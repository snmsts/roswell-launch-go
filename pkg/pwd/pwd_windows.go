// build windows

package pwd

import (
	"log"
	"strings"
	"syscall"
	"unsafe"
)

type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

var (
	FOLDERID_Profile = GUID{0x5E6C858F, 0x0E22, 0x4760, [8]byte{0x9A, 0xFE, 0xEA, 0x33, 0x17, 0xB6, 0x71, 0x73}}
)

var (
	modShell32               = syscall.NewLazyDLL("Shell32.dll")
	modOle32                 = syscall.NewLazyDLL("Ole32.dll")
	procSHGetKnownFolderPath = modShell32.NewProc("SHGetKnownFolderPath")
	procCoTaskMemFree        = modOle32.NewProc("CoTaskMemFree")
)

func SHGetKnownFolderPath(rfid *GUID, dwFlags uint32, hToken syscall.Handle, pszPath *uintptr) (retval error) {
	r0, _, _ := syscall.Syscall6(procSHGetKnownFolderPath.Addr(), 4, uintptr(unsafe.Pointer(rfid)), uintptr(dwFlags), uintptr(hToken), uintptr(unsafe.Pointer(pszPath)), 0, 0)
	if r0 != 0 {
		retval = syscall.Errno(r0)
	}
	return
}

func CoTaskMemFree(pv uintptr) {
	syscall.Syscall(procCoTaskMemFree.Addr(), 1, uintptr(pv), 0, 0)
	return
}

func ProfileFolder() (string, error) {
	var path uintptr
	err := SHGetKnownFolderPath(&FOLDERID_Profile, 0, 0, &path)
	if err != nil {
		return "", err
	}
	defer CoTaskMemFree(path)
	folder := syscall.UTF16ToString((*[1 << 16]uint16)(unsafe.Pointer(path))[:])
	str := strings.Replace(folder, "\\", "/", -1)
	return str, nil
}

func Systemhomedir() string {
	ret, err := ProfileFolder()
	if err != nil {
		log.Fatal(err)
	}
	return ret
}
