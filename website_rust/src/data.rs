use std::collections::HashMap;
use std::fs::File;
use std::io::{BufReader, Read};

use serde::Deserialize;

#[derive(Deserialize, Debug)]
pub struct Question {
    pub id: i32,
    pub difficulty: i32,
    pub title: String,
    pub slug: String,

    pub document: Option<String>,
    pub langs: Option<Vec<String>>,
    pub codes: Option<HashMap<String, String>>,

    pub prev_id: Option<i32>,
    pub next_id: Option<i32>,
}

impl Question {
    fn go_folder_name(&self) -> String {
        let mut folder_name = itoa::Buffer::new().format(self.id).to_owned();
        folder_name.push_str("_");
        folder_name.push_str(self.slug.replace("-", "_").as_str());
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
        file_path.push_str(itoa::Buffer::new().format(self.id));
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

    fn fill_data(&mut self) {
        self.document = Some(self.get_doc_content());

        let mut langs = vec![String::from("Go")];
        let mut codes = HashMap::new();
        codes.insert(String::from("Go"), self.get_go_content());
        if let Some(content) = self.get_rust_content() {
            langs.push(String::from("Rust"));
            codes.insert(String::from("Rust"), content);
        }
        self.langs = Some(langs);
        self.codes = Some(codes);

        match self.id {
            1 => self.next_id = Some(2),
            150 => self.prev_id = Some(149),
            _ => {
                self.prev_id = Some(self.id - 1);
                self.next_id = Some(self.id + 1);
            }
        }
    }
}

#[derive(Deserialize)]
pub struct Questions(Vec<Question>);

impl Questions {
    pub fn data(&self) -> &Vec<Question> {
        &self.0
    }

    pub fn find_by_id(&self, id: i32) -> Option<&Question> {
        for question in self.0.iter() {
            if question.id == id {
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
        let mut questions: Questions = serde_json::from_reader(reader).unwrap();

        for i in 0..questions.0.len() {
            questions.0[i].fill_data();
        }

        questions
    };
}
