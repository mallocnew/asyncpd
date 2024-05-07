// Copyright 2024 JOK Inc. All Rights Reserved.
// Author: easytojoin@163.com (jok)

#include <unistd.h>

#include <iostream>
#include <thread>  // NOLINT

void id_shower() {
  std::string thread_name;
  pthread_setname_np(pthread_self(), "ceshi");
  thread_name.resize(17);
  pthread_getname_np(pthread_self(), const_cast<char*>(thread_name.c_str()),
                     16);
  std::cout << thread_name << ": ";
  pthread_t tid = pthread_self();
  auto id = std::this_thread::get_id();
  std::cout << getpid() << " " << gettid() << " " << tid << " " << id
            << std::endl;
  while (true) {
  }
  return;
}

int main(int argc, char** argv) {
  pthread_setname_np(pthread_self(), "ceshi-main");
  char thread_name[100];
  pthread_getname_np(pthread_self(), thread_name, 100);
  std::cout << thread_name << ": ";
  pthread_t tid = pthread_self();
  auto id = std::this_thread::get_id();
  std::cout << getpid() << " " << gettid() << " " << tid << " " << id << "\n";
  std::thread task = std::thread(id_shower);
  std::cout << "new thread: " << task.get_id() << std::endl;
  task.join();
  std::cout << "thread: " << task.get_id() << " over" << std::endl;
  return 0;
}
