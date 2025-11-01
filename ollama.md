# Ollama
Secrets of Ollama: How it's built and the technologies used under the hood to power its operations.


1. Ollama <- `llama.cpp`
2. lmstudio.ai <- `llama.cpp`

Both use llama.cpp under the hood to run LLM models locally on the machine.


### Why ollama over lmstudio?
Take a look at the LM Studio GitHub repository: https://github.com/lmstudio-ai. It contains many tools, which can be confusing when trying to decide which one to use. Most of them have fewer GitHub stars compared to Ollama. Moreover, Ollama is developed in Go, while LM Studio is built with Python. Since Go is significantly faster and more efficient than Python, Ollama tends to run smoother and faster in practice.


### What is ineference?
to make `predictions` or `generate outputs` based on new input data.


### Ollama Model we will be using:

* ollama pull `qwen3:1.7b`
* ollama pull `llava:latest`
* ollama pull `mistral:latest`
* ollama pull `llama3.2:latest`
* ollama pull `nomic-embed-text`
* ollama pull `blaifa/Nanonets-OCR-s:latest`


## Learning Resource

* [LLM inference in C/C++](https://github.com/ggml-org/llama.cpp)
* https://lmstudio.ai/
