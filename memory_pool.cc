// Copyright 2024 JOK Inc. All Rights Reserved.
// Author: easytojoin@163.com (jok)

#include <iostream>
#include <list>
#include <mutex>  // NOLINT

template <typename T, std::size_t PoolSize = 100>
class MemoryPool {
 public:
  MemoryPool() { expandPool(); }
  ~MemoryPool() {
    std::unique_lock<std::mutex> lock(mutex_);
    for (auto& chunk : pool_) {
      delete[] reinterpret_cast<char*>(chunk);
    }
  }

  T* alloc() {
    std::unique_lock<std::mutex> lock(mutex_);
    if (free_chunks_.empty()) {
      expandPool();
    }
    T* ptr = free_chunks_.front();
    free_chunks_.pop_front();
    return ptr;
  }

  void dealloc(T* ptr) {
    std::unique_lock<std::mutex> lock(mutex_);
    free_chunks_.push_back(ptr);
  }

  std::size_t getFreeChunksCount() const {
    std::unique_lock<std::mutex> lock(mutex_);
    return free_chunks_.size();
  }

  std::size_t getUsedChunksCount() const {
    std::unique_lock<std::mutex> lock(mutex_);
    return PoolSize - getFreeChunksCount();
  }

 private:
  void expandPool() {
    char* blocks = new char[sizeof(T) * PoolSize];
    for (std::size_t i = 0; i < PoolSize; i++) {
      free_chunks_.push_back(reinterpret_cast<T*>(blocks + i * sizeof(T)));
    }
    pool_.push_back(blocks);
  }

 private:
  mutable std::mutex mutex_;
  std::list<T*> free_chunks_;
  std::list<char*> pool_;
};

struct DataObject {
  int data[100];
};

int main(int argc, char** argv) {
  MemoryPool<DataObject> pool;
  DataObject* obj1 = pool.alloc();
  DataObject* obj2 = pool.alloc();
  pool.dealloc(obj1);
  pool.dealloc(obj2);
  return 0;
}
