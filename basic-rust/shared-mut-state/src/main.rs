use std::sync::{Arc, Mutex};
use tokio::task;

#[tokio::main]
async fn main() {
    let counter = Arc::new(Mutex::new(0));

    let c1 = counter.clone();
    let t1 = task::spawn(async move {
        *c1.lock().unwrap() += 1;
    });

    let c2 = counter.clone();
    let t2 = task::spawn(async move {
        *c2.lock().unwrap() += 1;
    });

    t1.await.unwrap();
    t2.await.unwrap();
    
    println!("{}", *counter.lock().unwrap());
}
