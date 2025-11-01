# Whisper

> python3 --version

> python3 -m pip --version

upgrade pip
```
python3 -m pip install --upgrade pip
```

### Install ffmpeg

```
sudo apt remove --purge ffmpeg -y
sudo apt autoremove -y
```

```
sudo apt update
sudo apt install -y git build-essential pkg-config \
  nasm yasm libx264-dev libx265-dev libnuma-dev \
  libvpx-dev libfdk-aac-dev libmp3lame-dev libopus-dev \
  libass-dev libvorbis-dev libdav1d-dev libaom-dev \
  libbluray-dev libfreetype6-dev libfontconfig1-dev \
  libssl-dev
```

> `cd /usr/local/src`

```
sudo git clone https://github.com/FFmpeg/FFmpeg.git ffmpeg
cd ffmpeg
sudo git checkout n7.1.2
```

```
sudo ./configure \
  --prefix=/usr/local \
  --enable-gpl \
  --enable-nonfree \
  --enable-libx264 \
  --enable-libx265 \
  --enable-libvpx \
  --enable-libfdk-aac \
  --enable-libmp3lame \
  --enable-libopus \
  --enable-libass \
  --enable-libvorbis \
  --enable-libdav1d \
  --enable-libbluray \
  --enable-libfreetype \
  --enable-libfontconfig \
  --enable-openssl \
  --enable-libaom \
  --enable-pthreads \
  --enable-shared \
  --enable-version3
```

```
sudo make -j$(nproc)
sudo make install
```

```
echo "/usr/local/lib" | sudo tee /etc/ld.so.conf.d/ffmpeg.conf
sudo ldconfig
```

> `ffmpeg -version`


### Install Whisper
```
pip3 install git+https://github.com/openai/whisper.git
```

[whisper_downloading_state](./screens/whisper_downloading.png)

Folder structure

```
mkdir /root/aiagent
cd /root/aiagent
mkdir -p {input,output,done}
```
> python3 -m whisper --help

> recrod your voice using windows sound recorder
> `cp /mnt/c/Users/billr/OneDrive/เอกสาร/"Sound Recordings"/mostain_voice.wav .`

> `python3 -m whisper mostain_voice.wav --model medium`


location where models are downloaded:
* `ls -la ~/.cache/whisper/`
* `~/.cache/whisper/<model_name>.pt`

> `export HF_HOME=/opt/whisper-cache` `python3 -m whisper mostain_voice.wav --model medium`

> python3 -m whisper /root/aiagent/mostain_voice.wav --output_dir /root/aiagent/output --output_format srt --model medium

```
python3 -m whisper /root/aiagent/mostain_voice.wav \
  --output_dir /root/aiagent/output \
  --output_format srt \
  --model medium \
  --device cuda
```

> python3 -m whisper `/Users/.../input/{{$json.fileName}}` --output_dir `/Users/.../output` --output_format `all` --model `medium`

## Learning Resource
* https://docs.astral.sh/uv/
* https://www.vonage.com/
* https://ffmpeg.org/download.html
* https://github.com/openai/whisper
* https://huggingface.co/models?other=base_model:adapter:openai%2Fwhisper-medium&sort=trending
* https://huggingface.co/collections/openai/whisper-release-6501bba2cf999715fd953013
* https://huggingface.co/openai/whisper-large-v2
* https://chatgpt.com/c/68f49fe3-5884-8322-ae0b-beea6b1e00db