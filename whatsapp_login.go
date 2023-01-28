//gyd wong
package main

import (
    "fmt"
    "github.com/Rhymen/go-whatsapp"
)

func main() {
    // Buat sebuah session baru
    wac, err := whatsapp.NewConn(5 * time.Second)
    if err != nil {
        fmt.Println("Error saat membuat koneksi baru:", err)
        return
    }

    // Login ke akun WhatsApp
    if err := wac.Login(<nomor_telepon>, <kode_verifikasi>); err != nil {
        fmt.Println("Error saat login:", err)
        return
    }
    defer wac.Disconnect()

    fmt.Println("Berhasil login!")
}
