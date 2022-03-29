#[macro_use]
extern crate lazy_static;

use actix_web::{get, http::header, App, HttpResponse, HttpServer, Responder};
use askama::Template;

mod data;

#[derive(Template)]
#[template(path = "index.html")]
struct IndexTemplate<'a> {
    title: &'a str,
    questions: &'a Vec<data::Question>,
}

#[get("/")]
async fn index() -> impl Responder {
    HttpResponse::Ok()
        .append_header(header::ContentType::html())
        .body(
            IndexTemplate {
                title: "LeetCode 150",
                questions: data::QUESTIONS.data(),
            }
            .render()
            .unwrap(),
        )
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| App::new().service(index))
        .bind(("127.0.0.1", 15050))?
        .run()
        .await
}
