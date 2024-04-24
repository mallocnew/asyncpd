// Copyright 2024 GOTHAM Inc. All Rights Reserved.
// Author: easytojoin@163.com (jok)

use std::{
    sync::{mpsc, Arc, Mutex},
    thread,
};

fn first_thread() {
    let handler = thread::spawn(|| {
        println!("Hello world!");
        "Success"
    });
    let result = handler.join().unwrap();
    println!("{}", result);
}

fn counter() {
    let counter = Arc::new(Mutex::new(0));
    let handlers: Vec<_> = (0..10)
        .map(|_| {
            let counter = Arc::clone(&counter);
            thread::spawn(move || {
                let mut num = counter.lock().unwrap();
                *num += 1;
            })
        })
        .collect();
    for handler in handlers {
        handler.join().unwrap();
    }
    println!("Counter: {}", *counter.lock().unwrap());
}

fn communication() {
    let (sender, receiver) = mpsc::channel();
    thread::spawn(move || {
        sender.send(format!("Sender from {}", 1)).unwrap();
    });
    let message = receiver.recv().unwrap();
    println!("Received: {}", message);
}

fn main() {
    first_thread();
    counter();
    communication();
}
