use std::fs::File;
use std::io::{BufReader, Read};

use serde::Deserialize;

#[derive(Deserialize)]
pub struct QuestionData {
    pub id: i32,
    pub difficulty: i32,
    pub title: String,
    pub slug: String,
}

pub struct Question {
    pub data: QuestionData,

    pub document: String,
    pub langs: Vec<String>,
    pub display_langs: Vec<String>,
    pub codes: Vec<String>,

    pub prev_id: Option<i32>,
    pub next_id: Option<i32>,
}

impl Question {
    fn go_folder_name(&self) -> String {
        let mut folder_name = itoa::Buffer::new().format(self.data.id).to_owned();
        folder_name.push_str("_");
        folder_name.push_str(self.data.slug.replace("-", "_").as_str());
        folder_name
    }

    fn doc_file_path(&self) -> String {
        let mut file_path = String::from("../documents/");
        file_path.push_str(self.go_folder_name().as_str());
        file_path.push_str(".md");
        file_path
    }

    fn get_doc_content(&self) -> String {
        let mut file = File::open(self.doc_file_path().as_str()).unwrap();
        let mut content = String::new();
        file.read_to_string(&mut content);
        markdown::to_html(content.as_str())
    }

    fn get_go_content(&self) -> String {
        let mut file_path = String::from("../go/");
        file_path.push_str(self.go_folder_name().as_str());
        file_path.push_str("/main.go");
        let mut file = File::open(file_path.as_str()).unwrap();
        let mut content = String::new();
        file.read_to_string(&mut content);

        let strs: Vec<&str> = content.split("// code\n").collect();
        strs[1].to_owned()
    }

    fn rust_file_path(&self) -> String {
        let mut file_path = String::from("../rust/src/question_");
        file_path.push_str(itoa::Buffer::new().format(self.data.id));
        file_path.push_str(".rs");
        file_path
    }

    fn get_rust_content(&self) -> Option<String> {
        let file_path = self.rust_file_path();
        let mut file = File::open(file_path.as_str());
        if file.is_err() {
            return None;
        }

        let mut file = File::open(file_path.as_str()).unwrap();
        let mut content = String::new();
        file.read_to_string(&mut content);

        let mut strs: Vec<&str> = content.split("struct Solution;\n\n").collect();
        strs = strs[1].split("\n#[test]").collect();
        Some(strs[0].to_owned())
    }

    fn fill_info(&mut self) {
        self.document = self.get_doc_content();

        self.langs.push(String::from("go"));
        self.display_langs.push(String::from("Go"));
        self.codes.push(self.get_go_content());
        if let Some(content) = self.get_rust_content() {
            self.langs.push(String::from("rust"));
            self.display_langs.push(String::from("Rust"));
            self.codes.push(content);
        }

        match self.data.id {
            1 => self.next_id = Some(2),
            150 => self.prev_id = Some(149),
            _ => {
                self.prev_id = Some(self.data.id - 1);
                self.next_id = Some(self.data.id + 1);
            }
        }
    }
}

pub struct Questions(Vec<Question>);

impl Questions {
    pub fn data(&self) -> &Vec<Question> {
        &self.0
    }

    pub fn find_by_id(&self, id: i32) -> Option<&Question> {
        for question in self.0.iter() {
            if question.data.id == id {
                return Some(question);
            }
        }
        None
    }
}

lazy_static! {
    pub static ref QUESTIONS: Questions = {
        let file = File::open("src/data.json").unwrap();
        let reader = BufReader::new(file);
        let mut question_data: Vec<QuestionData> = serde_json::from_reader(reader).unwrap();

        let mut questions = Vec::new();
        for data in question_data {
            let mut question = Question {
                data: data,

                document: String::from(""),
                langs: Vec::new(),
                display_langs: Vec::new(),
                codes: Vec::new(),

                prev_id: None,
                next_id: None,
            };
            question.fill_info();
            questions.push(question);
        }

        Questions(questions)
    };
}
