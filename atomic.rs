// Copyright 2024 JOK Inc. All Rights Reserved.
// Author: easytojoin@163.com (jok)

use std::{
    sync::{atomic::AtomicBool, Arc},
    thread,
    time::Duration,
};

fn main() {
    let flag = Arc::new(AtomicBool::new(false));
    let c = Arc::clone(&flag);
    let handler = thread::spawn(move || {
        while !c.load(std::sync::atomic::Ordering::Relaxed) {
            thread::sleep(Duration::from_secs(1));
            println!("Thread working...");
        }
        println!("Thread over");
    });
    thread::sleep(Duration::from_secs(10));
    flag.store(true, std::sync::atomic::Ordering::Relaxed);
    handler.join().unwrap();
}
