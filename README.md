# black-vault-cli

CLI สำหรับ **Git BlackVault** (Go + Cobra) — ใช้ร่วมกับ [black-vault-lib](../black-vault-lib) และเชื่อมกับ [black-vault-gui](../black-vault-gui) ผ่าน gRPC

## สรุป

- **ภาษา:** Go
- **คำสั่ง:** `open`, `close`, `status`, `serve`
- **ไลบรารี:** เรียกใช้ [black-vault-lib](../black-vault-lib) สำหรับ config, workspace, clone, **SQLite store** (registry + cache)
- **ข้อมูล/cache:** เก็บที่ lib — ไฟล์ `~/.blackvault/blackvault.db` (สร้างใหม่ถ้าไม่มี)
- **gRPC:** คำสั่ง `serve` รัน gRPC server ให้ GUI เชื่อมต่อ (ต้อง generate proto ก่อน)

## โครงสร้าง

```
black-vault-cli/
├── api/proto/
│   └── blackvault.proto   # สัญญา gRPC (ซิงค์กับ black-vault-lib)
├── cmd/
│   ├── root.go
│   ├── open.go
│   ├── close.go
│   ├── status.go
│   └── serve.go
├── main.go
├── go.mod                 # replace black-vault-lib => ../black-vault-lib
├── Makefile
└── README.md
```

## สิ่งที่ทำไปแล้ว

- **Cobra root** พร้อมคำสั่งย่อย
- **open [group/repo]** — clone repo (ใช้ system git หรือ portable ตาม config)
- **close [group/repo]** — ลบ workspace (มี `--force`)
- **status** — แสดงรายการ ACTIVE (และข้อความ CLOSED)
- **config [get|set]** — ดู/ตั้งค่า config เช่น **git_path** (ใช้ git จากไหน: ไม่ตั้ง = ใช้ของระบบเป็น default, ไม่มีถึงค่อยใช้ portable)
- **install-git** — สร้างโฟลเดอร์สำหรับ portable git (`~/.blackvault/tools/git`) และแสดงวิธีดาวน์โหลด/แตก แล้วตั้ง `git_path` ได้
- **serve** — รัน gRPC server; **dynamic port**: ถ้า `--port` (default 50051) ถูกใช้ จะลอง port ถัดไป (50052, 50053, …) แล้วเขียนพอร์ตที่ใช้ลง `~/.blackvault/grpc_port` ให้ GUI อ่าน
- **Proto** — เก็บ `api/proto/blackvault.proto` ไว้ใน repo (ซิงค์กับ lib) สำหรับ generate Go/Dart ฝั่ง CLI และ GUI

## การ build

```bash
# จากโฟลเดอร์ black-vault-cli
go build -o blackvault .
# หรือ
make build
```

## Generate gRPC (สำหรับ serve + GUI)

ต้องมี `protoc` และ plugins:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# ติดตั้ง protoc (เช่น brew install protobuf บน macOS)
make proto
```

จากนั้นแก้ `cmd/serve.go` ให้ register BlackVaultService และใช้ generated code แล้ว build ใหม่

## การใช้

```bash
./blackvault open group/repo
./blackvault open group/repo --shallow
./blackvault close group/repo
./blackvault status
./blackvault config get
./blackvault config set git_path /path/to/git
./blackvault install-git
./blackvault serve --port 50051
```

- **Git setting:** ถ้าไม่ตั้ง `git_path` จะใช้ **git ของระบบ (PATH)** เป็น default ถ้าไม่มีจะดูที่ `~/.blackvault/tools/git` (portable). ตั้งเอง: `blackvault config set git_path <path>` หรือรัน `blackvault install-git` แล้วเอา portable ไปวางแล้วตั้ง path
- **Dynamic port:** ถ้า port ชน จะลอง port ถัดไปอัตโนมัติ และเขียนพอร์ตที่ใช้ลง `~/.blackvault/grpc_port` (ลบเมื่อ serve ปิด) — GUI อ่านไฟล์นี้เพื่อเชื่อมต่อ

Config อยู่ที่ `~/.blackvault/config.yaml` (สร้างโดยอัตโนมัติเมื่อไม่มี หรือใช้จาก lib).
