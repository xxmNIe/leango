package dockerCmd

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"com.py/common"
)

/**
mark
100 downloading
110 success
120 image is not found or download failed
*/

//此代码主要实现Docker的下载、创建、删除和查询操作

//由于Docker的数据无法迁移，因此并不推荐使用Docker提供有状态的服务
//DownlodaContainer 用指定名称下载Docker镜像

type Owner struct {
	imageName   string `json:"imageName"`
	creatorName string `json:"creatorName"`
}

type Downloadinfo struct {
	imageName string
}

const (
	IMAGE_DOWNLOADING = iota + 1
	IAMGE_DOWNSUCCESS
	IAMGE_DOWNFAILED
)

//var mark map[Owner]int
var mark map[*Downloadinfo]int
var markLock sync.RWMutex

func DownloadImage(imageName string, creatorName string) (msg string, ow Owner) {
	// ow.imageName = imageName
	// ow.creatorName = creatorName
	// markLock.Lock()
	// mark[ow] = 1
	// markLock.Unlock()
	// go func() {
	// 	if output, _ := exec.Command("docker", "images").CombinedOutput(); strings.Contains(string(output), ow.imageName) {
	// 		markLock.Lock()
	// 		mark[ow] = 2
	// 		markLock.Unlock()
	// 	}
	// 	cmd := exec.Command("docker", "pull", ow.imageName)
	// 	_, err := cmd.CombinedOutput()
	// 	//println("GO ROUTINE")
	// 	//println("Output",string(output))
	// 	//println("Error",err)
	// 	//_, err := file.WriteString(a)
	// 	markLock.Lock()
	// 	defer markLock.Unlock()
	// 	if err == nil {
	// 		mark[ow] = 2
	// 		//println("SUCCESS mark[a]:", mark[imageName])
	// 	} else {
	// 		mark[ow] = -1
	// 		println("FAILED mark[a]", mark[ow])
	// 	}
	// }()
	// return "DOWNLOADING", ow
	ow.imageName = imageName
	ow.creatorName = creatorName
	downLoad := &Downloadinfo{
		imageName: imageName,
	}
	// markLock.RLock()
	// if status, ok := mark[downLoad]; ok && status != 3 {
	// 	markLock.RUnlock()
	// 	return "DOWNLOADING", ow
	// }
	// markLock.RUnlock()
	markLock.Lock()
	if status, ok := mark[downLoad]; ok {
		if status == 3 {
			delete(mark, downLoad)
		}
		markLock.Unlock()
		return "DOWNLOADING", ow
	}
	mark[downLoad] = 1
	markLock.Unlock()

	go func() {
		output, err := exec.Command("dockerg", "images").CombinedOutput()
		if err != nil {
			markLock.Lock()
			mark[downLoad] = 3
			markLock.Unlock()
			return
		}
		if strings.Contains(string(output), downLoad.imageName) {
			markLock.Lock()
			mark[downLoad] = 2
			markLock.Unlock()
			return
		}

		cmd := exec.Command("docker", "pull", downLoad.imageName)
		_, err = cmd.CombinedOutput()
		markLock.Lock()
		defer markLock.Unlock()
		if err == nil {
			mark[downLoad] = 2
			//println("SUCCESS mark[a]:", mark[imageName])
		} else {
			mark[downLoad] = 3
			println("FAILED mark[a]", mark[downLoad])
		}

	}()
	return "DOWNLOADING", ow

}

func CheckDownloading(own Owner) (msg string, code int) {

	// markLock.Lock()
	// defer markLock.Unlock()
	// if mark[own] == 1 {
	// 	//println("DOWNLOADING CHECK", mark[a])
	// 	return "镜像准备中", common.Status_ImageDownloading
	// } else if mark[own] == 2 {
	// 	println("SUCCESS CHECK", mark[own])
	// 	delete(mark, own)
	// 	return "镜像准备完成", common.Status_ImageDownload_success
	// } else {
	// 	delete(mark, own)
	// 	return "镜像下载失败", common.Status_ImageDownload_err
	// }
	downInfo := &Downloadinfo{imageName: own.imageName}
	markLock.RLock()
	defer markLock.Unlock()
	if mark[downInfo] == 1 {
		//println("DOWNLOADING CHECK", mark[a])
		return "镜像准备中", common.Status_ImageDownloading
	} else if mark[downInfo] == 2 {
		println("SUCCESS CHECK", mark[downInfo])
		delete(mark, downInfo)
		return "镜像准备完成", common.Status_ImageDownload_success
	} else {
		return "镜像下载失败", common.Status_ImageDownload_err
	}

}

func RemoveImage(imageName string) error {
	//cwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	downloadImageCMD := exec.Command("docker", "rmi", imageName)
	err := downloadImageCMD.Run()
	return err
}

//CreateContainer 用指定名称创建Docker容器
func CreateContainer(containerName, imageName, path string) error {
	cwd, _ := filepath.Abs(filepath.Dir(path))
	createContainerCMD := exec.Command("docker", "run", "-id", "--name="+containerName, "--net=none", "-v", cwd+"/docker/"+containerName+"/sandbox:/root", imageName)
	err := createContainerCMD.Run()
	return err
}

//KillAndRemoveContainer 杀死并删除一个Docker容器
func KillAndRemoveContainer(containerName string) {
	cmd := exec.Command("docker", "kill", containerName)
	cmd.Run()
	cmd = exec.Command("docker", "rm", containerName)
	cmd.Run()
	os.RemoveAll("./docker/" + containerName)
}

func SearchImages(imageName string) bool {
	cmd := exec.Command("docker", "images")
	res, _ := cmd.CombinedOutput()
	flag := strings.Contains(string(res), imageName)
	return flag
}
