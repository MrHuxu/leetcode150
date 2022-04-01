#[macro_use]
extern crate lazy_static;

use actix_web::{
    get, http::header, middleware::Logger, web, App, HttpResponse, HttpServer, Responder,
};
use askama::Template;

mod data;

#[derive(Template)]
#[template(path = "error.html")]
struct ErrorTemplate<'a> {
    title: &'a str,
    meta: &'a str,
    message: &'a str,
}

#[derive(Template)]
#[template(path = "index.html")]
struct IndexTemplate<'a> {
    title: &'a str,
    meta: &'a str,
    questions: &'a Vec<data::Question>,
}

#[get("/")]
async fn index() -> impl Responder {
    HttpResponse::Ok()
        .append_header(header::ContentType::html())
        .body(
            IndexTemplate {
                title: "LeetCode 150 - xhu",
                meta: "LeetCode 150 - xhu",
                questions: data::QUESTIONS.data(),
            }
            .render()
            .unwrap(),
        )
}

#[derive(Template)]
#[template(path = "detail.html")]
struct DetailTemplate<'a> {
    title: &'a str,
    meta: &'a str,
    question: &'a data::Question,
}

#[get("/{id}")]
async fn detail(params: web::Path<i32>) -> impl Responder {
    if let Some(question) = data::QUESTIONS.find_by_id(params.into_inner()) {
        let mut title = itoa::Buffer::new().format(question.data.id).to_owned();
        title.push_str(". ");
        title.push_str(question.data.title.as_str());
        title.push_str(" - xhu");

        HttpResponse::Ok()
            .append_header(header::ContentType::html())
            .body(
                DetailTemplate {
                    title: title.as_str(),
                    meta: question.document.as_str(),
                    question: question,
                }
                .render()
                .unwrap(),
            )
    } else {
        HttpResponse::Ok()
            .append_header(header::ContentType::html())
            .body(
                ErrorTemplate {
                    title: "Server Error",
                    meta: "",
                    message: "question not found",
                }
                .render()
                .unwrap(),
            )
    }
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .wrap(Logger::default())
            .service(index)
            .service(detail)
    })
    .bind(("0.0.0.0", 15050))?
    .run()
    .await
}
