use axum::extract::Path;
use axum::{
    http::StatusCode,
    response::{IntoResponse, Result},
    routing::get,
    Json, Router,
};
use serde::{Deserialize, Serialize};

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct User {
    pub id: i64,
    pub name: String,
    pub username: String,
    pub email: String,
    pub address: Address,
    pub phone: String,
    pub website: String,
    pub company: Company,
}

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct Address {
    pub street: String,
    pub suite: String,
    pub city: String,
    pub zipcode: String,
    pub geo: Geo,
}

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct Geo {
    pub lat: String,
    pub lng: String,
}

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct Company {
    pub name: String,
    pub catch_phrase: String,
    pub bs: String,
}

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct Todo {
    pub user_id: i32,
    pub id: i32,
    pub title: String,
    pub completed: bool,
}

async fn user_handler() -> Result<impl IntoResponse, (StatusCode, Json<serde_json::Value>)> {
    let users = match get_users().await {
        Ok(users) => serde_json::json!(users),
        Err(e) => {
            let error_response: serde_json::Value = serde_json::json!({"status": "INTERNAL_SERVER_ERROR","message": format!("{:?}", e)});
            return Err((StatusCode::INTERNAL_SERVER_ERROR, Json(error_response)));
        }
    };

    Ok(Json(users))
}

async fn todo_handler(
    Path(id): Path<i32>,
) -> Result<impl IntoResponse, (StatusCode, Json<serde_json::Value>)> {
    let todos = match get_todos_by_user_id(id).await {
        Ok(todos) => serde_json::json!(todos),
        Err(e) => {
            let error_response: serde_json::Value = serde_json::json!({"status": "INTERNAL_SERVER_ERROR","message": format!("{:?}", e)});
            return Err((StatusCode::INTERNAL_SERVER_ERROR, Json(error_response)));
        }
    };

    Ok(Json(todos))
}

async fn get_users() -> Result<Vec<User>, Box<dyn std::error::Error>> {
    let user_url: &'static str = env!("USER_URL");
    let users: Vec<User> = reqwest::get(format!("{user_url}/users"))
        .await?
        .json()
        .await?;
    Ok(users)
}

async fn get_todos_by_user_id(user_id: i32) -> Result<Vec<Todo>, Box<dyn std::error::Error>> {
    let todo_url: &'static str = env!("TODO_URL");
    let todos: Vec<Todo> = reqwest::get(format!("{todo_url}/todos"))
        .await?
        .json()
        .await?;

    let result = todos
        .into_iter()
        .filter(|voc| voc.user_id == user_id)
        .collect::<Vec<Todo>>();
    Ok(result)
}

#[tokio::main]
async fn main() {
    // Define Routes
    let app = Router::new()
        .route("/users", get(user_handler))
        .route("/users/:id/todos", get(todo_handler));

    println!("Running on http://localhost:3000");
    // Start Server
    axum::Server::bind(&"127.0.0.1:3000".parse().unwrap())
        .serve(app.into_make_service())
        .await
        .unwrap();
}
