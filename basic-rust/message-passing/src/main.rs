use tokio::sync::mpsc;

#[tokio::main]
async fn main() {
    let (tx, mut rx) = mpsc::channel(1);

    tokio::spawn(async move {
        tx.send(10).await.unwrap();
    });
    
    let v = rx.recv().await.unwrap();
    println!("{}", v);
}
