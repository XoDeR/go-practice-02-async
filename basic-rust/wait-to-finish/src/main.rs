async fn worker() ->i32 {
    345
}

#[tokio::main]
async fn main() {
    let handle = tokio::spawn(worker());
    let result = handle.await.unwrap();
    
    println!("{}", result);
}
