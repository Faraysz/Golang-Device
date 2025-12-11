package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// Struct = mirip seperti model di Laravel / TypeScript
type Device struct {
    UID   string
    Name  string
    Status string
}

var devices []Device

// Fungsi untuk tambah device
func addDevice() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Masukkan UID Device: ")
    uid, _ := reader.ReadString('\n')

    fmt.Print("Masukkan Nama Device: ")
    name, _ := reader.ReadString('\n')

    devices = append(devices, Device{
        UID:   strings.TrimSpace(uid),
        Name:  strings.TrimSpace(name),
        Status: "online",
    })

    fmt.Println("Device berhasil ditambahkan!")
}

// Fungsi untuk menampilkan semua device
func listDevices() {
    fmt.Println("\nDaftar Device:")
    for _, d := range devices {
        fmt.Printf("- %s (%s) status: %s\n", d.Name, d.UID, d.Status)
    }
    fmt.Println()
}

// Fungsi cari device by UID
func findDevice() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Masukkan UID yang ingin dicari: ")
    uid, _ := reader.ReadString('\n')
    uid = strings.TrimSpace(uid)

    for _, d := range devices {
        if d.UID == uid {
            fmt.Printf("Device ditemukan: %s (%s) status: %s\n", d.Name, d.UID, d.Status)
            return
        }
    }

    fmt.Println("Device tidak ditemukan.")
}

// MAIN PROGRAM (titik masuk golang)
func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Println("\n=== MENU DEVICE MANAGER ===")
        fmt.Println("1. Tambah Device")
        fmt.Println("2. Lihat Semua Device")
        fmt.Println("3. Cari Device by UID")
        fmt.Println("4. Keluar")
        fmt.Print("Pilih menu: ")

        scanner.Scan()
        choice := scanner.Text()

        switch choice {
        case "1":
            addDevice()
        case "2":
            listDevices()
        case "3":
            findDevice()
        case "4":
            fmt.Println("Keluar program...")
            os.Exit(0)
        default:
            fmt.Println("Menu tidak valid, coba lagi!")
        }
    }
}
