#[tokio::main]
async fn main() {
    tokio::spawn(async {
        println!("From task");
    });

    tokio::time::sleep(std::time::Duration::from_secs(1)).await;
}
