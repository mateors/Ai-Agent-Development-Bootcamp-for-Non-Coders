

You can convert an image to a base64-encoded string using several methods depending on your environment. Here are a few options:

---

### 🖥️ **Using the Command Line (Linux/macOS)**
```bash
base64 image.png > output.txt
```
- This reads `image.png` and writes the base64 string to `output.txt`.

To get the string directly:
```bash
base64 image.png
```

---

### 🧑‍💻 **Using Python**
```python
import base64

with open("image.png", "rb") as img_file:
    b64_string = base64.b64encode(img_file.read()).decode('utf-8')
    print(b64_string)
```

---

### 🌐 **Using JavaScript (Browser)**
```javascript
const fileInput = document.querySelector('input[type="file"]');
fileInput.addEventListener('change', () => {
  const reader = new FileReader();
  reader.onload = () => {
    const base64String = reader.result.split(',')[1];
    console.log(base64String);
  };
  reader.readAsDataURL(fileInput.files[0]);
});
```

---

### 📦 **Using Go**
Since you're a Go developer, here's a Go snippet:
```go
package main

import (
    "encoding/base64"
    "fmt"
    "io/ioutil"
)

func main() {
    data, err := ioutil.ReadFile("image.png")
    if err != nil {
        panic(err)
    }
    encoded := base64.StdEncoding.EncodeToString(data)
    fmt.Println(encoded)
}
```

Here’s a clean one‑liner that will take an image file, encode it to base64, and inject it directly into your curl request

```
curl http://localhost:11434/api/chat \
  -H "Content-Type: application/json" \
  -d '{
    "model": "llama3.2-vision",
    "messages": [
      {
        "role": "user",
        "content": "what is in this image?",
        "images": ["'"$(base64 -i image.png | tr -d '\n')"'"]
      }
    ]
  }'
```
