# GPU

### Check GPU status
> nvidia-smi

> nvcc --version

### Monitor GPU usage in real time
> watch -n 1 nvidia-smi


> `python3 -m torch.utils.collect_env`

### Check if PyTorch detects your GPU:
```
python3 -c "import torch; print(torch.cuda.is_available(), torch.cuda.get_device_name(0))"
```

```
python3 -m whisper /root/aiagent/mostain_voice.wav \
  --output_dir /root/aiagent/output \
  --output_format srt \
  --model medium \
  --device cuda
```

### If you want to confirm PyTorch was compiled with CUDA, run:
```
python3 -c "import torch; print(torch.version.cuda)"
```