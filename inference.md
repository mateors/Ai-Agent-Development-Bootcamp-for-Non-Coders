# Inference
to make predictions or generate outputs

Inference in the context of machine learning — especially large language models (LLMs) — refers to the process of **using a trained model to make predictions or generate outputs** based on new input data.

---

### 🧠 Inference vs Training

| Phase      | Purpose                          | Example                        |
|------------|----------------------------------|--------------------------------|
| **Training** | Learn patterns from labeled data | Teaching a model grammar rules |
| **Inference** | Apply learned patterns to new data | Asking the model to write a poem |

During inference, the model:
- Loads its trained weights
- Accepts input (e.g., a prompt, image, or question)
- Runs forward computations (no gradient updates)
- Produces output (e.g., text, classification, translation)

---

### 🔍 Examples of Inference

- You type: `"Translate 'Hello' to French"` → Model infers: `"Bonjour"`
- You upload an image → OCR model infers: `"Signature detected: MD Billah"`
- You ask: `"What’s the capital of Japan?"` → Model infers: `"Tokyo"`

---

### ⚙️ Inference in LLMs

In LLMs like GPT, LLaMA, or Mixtral:
- Inference involves **token-by-token generation**
- It uses **attention mechanisms** and **transformer layers**
- It’s optimized for speed and low latency — especially when quantized