#[tokio::main]
async fn main() {
    let body = reqwest::get("https://example.com")
        .await
        .unwrap()
        .text()
        .await
        .unwrap();

    println!("{}", body);
}
