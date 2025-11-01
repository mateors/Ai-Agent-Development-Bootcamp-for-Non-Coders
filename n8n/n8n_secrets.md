# n8n secrets

lesser-known but powerful **n8n secrets** that can supercharge your workflows and make your automation game modular, panic-proof, and scalable—just the way you like it:

---

### 🧠 1. **You Can Use JavaScript Everywhere**
Expressions in n8n support full JavaScript, not just simple variables. That means you can do things like:

```js
{{ $json.amount > 1000 ? "High Value" : "Standard" }}
```

Or even:

```js
{{ new Date().toISOString().split("T")[0] }}
```

Use this to dynamically format dates, conditionally route logic, or build custom strings.

---

### 🧩 2. **Hidden `Run Once for All Items` Toggle**
In nodes like “Set” or “Function,” there's a hidden toggle under the gear icon: **“Run Once for All Items.”**  
This lets you process all items together—perfect for aggregations, summaries, or batch logic.

---

### 🧵 3. **You Can Create Loops Without the Loop Node**
Use a combination of:
- **Webhook → Function → HTTP Request → Respond to Webhook**
- Or **Execute Workflow → Wait → Merge**

This lets you build recursive or conditional loops, especially useful for paginated APIs or retry logic.

---

### 🧪 4. **Use `IF` Node as a Switch**
The `IF` node isn’t just for true/false. You can chain multiple `IF` nodes to simulate a **switch-case** structure.  
Bonus: You can use expressions like:

```js
{{ $json.status === "pending" }}
```

---

### 🧰 5. **Access Execution Metadata with `$execution`**
You can log or return metadata like:

```js
{
  "workflowId": $workflow.id,
  "executionId": $execution.id,
  "startedAt": $execution.startedAt
}
```

Perfect for audit trails or debugging.

---

### 🧼 6. **Clean Up with `Empty` Node**
Want to remove fields from your data? Use the **“Set” node** with `Keep Only Set` enabled and leave it blank.  
It acts like a **data sanitizer**, stripping everything.

---

### 🧠 7. **Use `Function` Node to Create Custom JSON Structures**
Instead of struggling with the “Set” node, use a `Function` node:

```js
return [
  {
    json: {
      invoice: {
        items: $items.map(item => item.json),
        total: $items.reduce((sum, item) => sum + item.json.amount, 0)
      }
    }
  }
];
```

This gives you full control over nesting and formatting.

---

### 🧙‍♂️ 8. **You Can Trigger Workflows via Internal HTTP**
Every workflow with a webhook can be triggered internally using `http://localhost:5678/webhook/...`  
Useful for chaining workflows without exposing them externally.

---

### 🧭 9. **Use `Wait` Node for Human-in-the-Loop**
Want to pause until someone approves something? Use the **“Wait” node** with a webhook to resume.  
This enables **human-in-the-loop automation**—great for approvals, reviews, or manual overrides.

---

### 🧬 10. **You Can Modularize with `Execute Workflow`**
Split your logic into reusable workflows and call them using the **“Execute Workflow”** node.  
This keeps your main workflow clean and lets you reuse logic across projects.