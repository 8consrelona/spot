# Faceswap Dev
  
so ive been trying to use faceswap to make funny deepfake videos i have a ryzen cpu and an amd rx6800m dedicated GPU. my problem is that opensuse tumbleweed is not supported with the amd pro gpu driver for opencl witch is required for faceswap plaidml so that option is out and the other amd option os the ROCm support option wich is also not supported in tumbleweed what can i do to make deepfakes or get faceswap to work with my distro?
 
note i also tried to use kvm justto pass my gpu to windows vm but it did show up in windows but the driver would alwayse crash in the vm so idk whhat i did wrong there a tutortial i followed said i had to blacklist my dgpu and i really dont wanna do that then change it all back in terminals and with nano every f\*g time i wanna play a game and or use my vm. not unless there is something im missing i have iommu or whatever it is in bios enabled i used a ueffi bios firmware in the kvm so yea and i added the audio controller and the dgpu to the kvm before booting it.
 
**DOWNLOAD ••• [https://comlum-profhe.blogspot.com/?augy=2A0Tbr](https://comlum-profhe.blogspot.com/?augy=2A0Tbr)**


 
I would suggest just stick to proton/wine for graphics-heavy gaming, windows vm without gpu passthrough (install this when you need opengl: GitHub - pal1000/mesa-dist-win: Pre-built Mesa3D drivers for Windows) can be used for light gaming. If you really want windows, dual-booting (or with wsl) is much easier.
 
ok thanks i tried i was hoping for an easy or well at least a solution that didnt involve me hving to reboot but i just installed a copy of windows 10 to a usb drive now its portable and works fine but thanks anyway ^^
 
I have tried to setup faceswap.py peviously on an old pc and it worked.I have chosen the same configuration and it somehow does not work. When I try to open it with the power shell, it gives me this error;
 
FaceSwapLab is an extension for Stable Diffusion that simplifies the use of insighface models for face-swapping. It has evolved from sd-webui-faceswap and some part of sd-webui-roop. However, a substantial amount of the code has been rewritten to improve performance and to better manage masks.
 
Some key features include the ability to reuse faces via checkpoints, multiple face units, batch process images, sort faces based on size or gender, and support for vladmantic. It also provides a face inpainting feature.
 
This extension is **not intended to facilitate the creation of not safe for work (NSFW) or non-consensual deepfake content**. Its purpose is to bring consistency to image creation, making it easier to repair existing images, or bring characters back to life.
 
We will comply with European regulations regarding this type of software. As required by law, the code may include both visible and invisible watermarks. If your local laws prohibit the use of this extension, you should not use it.

 a2f82b0cb4
 
