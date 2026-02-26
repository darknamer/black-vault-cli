# black-vault-cli

CLI สำหรับ **Git BlackVault** (Go + Cobra) — ใช้ร่วมกับ [black-vault-lib](../black-vault-lib) และเชื่อมกับ [black-vault-gui](../black-vault-gui) ผ่าน gRPC

## สรุป

- **ภาษา:** Go (คอมเมนต์ในโค้ดเป็นภาษาไทย)
- **คำสั่ง:** `open`, `close`, `status`, `serve`, และคำสั่ง Git (commit, branch, remote, git-flow)
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
│   ├── serve.go
│   ├── git_commit.go    # commit, คอมเมนต์ภาษาไทย
│   ├── git_add.go       # git add (paths... หรือ -A)
│   ├── git_fetch.go     # git fetch (remote/refspec หรือ --all/--prune)
│   ├── git_pull.go      # git pull (remote/branch, --rebase)
│   ├── git_push.go      # git push (remote/branch, -u, --force)
│   ├── git_branch.go    # สร้าง/สลับ/เปลี่ยนชื่อ/ลบ branch, ตั้ง upstream
│   ├── git_merge.go     # รวม branch (git merge)
│   ├── git_remote.go    # จัดการ remote (list/add/remove/set-url)
│   └── git_flow.go      # git-flow feature/release/hotfix
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
- **คำสั่ง Git (คอมเมนต์ภาษาไทย):**
  - `git-commit [group/repo] -m "ข้อความ"` — commit (ดีฟอลต์ทำ `git add -A` ก่อน), ใช้ `--no-add` ถ้าไม่ต้องการ add
  - `git-add [group/repo] [paths...]` — รัน git add; ถ้าไม่ระบุ path จะใช้ `git add -A` (เพิ่มทั้งหมด)
  - `git-fetch [group/repo]` — รัน git fetch (รองรับ `--all`, `--prune`, `--remote/-r`, `--refspec`)
  - `git-pull [group/repo]` — รัน git pull (remote/branch ตาม flag, รองรับ `--rebase`)
  - `git-push [group/repo]` — รัน git push (remote/branch ตาม flag, รองรับ `-u/--set-upstream`, `--force`)
  - `git-branch-create [group/repo] [branch]` — สร้าง branch ใหม่และ checkout
  - `git-branch-switch [group/repo] [branch]` — สลับไป branch ที่มีอยู่ (checkout)
  - `git-branch-rename [group/repo] [new_name]` — เปลี่ยนชื่อ branch ปัจจุบัน
  - `git-branch-rename [group/repo] [old_name] [new_name]` — เปลี่ยนชื่อจาก old เป็น new
  - `git-branch-delete [group/repo] [branch]` — ลบ branch ในพื้นที่, ใช้ `-f/--force` เพื่อบังคับลบ (-D)
  - `git-branch-set-upstream [group/repo] [upstream] [branch]` — ตั้ง upstream ของ branch (เช่น origin/main), ไม่ระบุ branch = branch ปัจจุบัน
  - `git-merge [group/repo] [branch]` — รวม branch ที่ระบุเข้า branch ปัจจุบัน (`--no-ff`, `--squash`, `-m` ได้)
  - `git-remote list [group/repo]` — แสดงรายการ remote
  - `git-remote add [group/repo] [name] [url]` — เพิ่ม remote
  - `git-remote remove [group/repo] [name]` — ลบ remote
  - `git-remote set-url [group/repo] [name] [url]` — เปลี่ยน URL ของ remote
  - `git-flow init [group/repo]` — เตรียมโครง git-flow แบบง่าย (main/master + develop)
  - `git-flow feature-start [group/repo] [name]` — เริ่ม feature branch: feature/<name> แล้ว checkout
  - `git-flow release-start [group/repo] [name]` — เริ่ม release branch: release/<name> จาก develop (หรือ branch ที่กำหนด)
  - `git-flow release-finish [group/repo] [name]` — จบ release: merge เข้า main + develop (มีตัวเลือกสร้าง tag)
  - `git-flow hotfix-start [group/repo] [name]` — เริ่ม hotfix branch: hotfix/<name> จาก main (หรือ branch ที่กำหนด)
  - `git-flow hotfix-finish [group/repo] [name]` — จบ hotfix: merge เข้า main + develop (มีตัวเลือกสร้าง tag)

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
# Workspace
./blackvault open group/repo
./blackvault open group/repo --shallow
./blackvault close group/repo
./blackvault status

# Git: add/commit, fetch/pull/push, branch, merge, remote, git-flow (รายละเอียดดูด้านบน)
./blackvault git-commit group/repo -m "ข้อความ"
./blackvault git-add group/repo              # เทียบเท่า git add -A
./blackvault git-add group/repo path/to/file.go other/file.txt
./blackvault git-fetch group/repo --all --prune
./blackvault git-fetch group/repo -r origin --refspec 'refs/heads/*:refs/remotes/origin/*'
./blackvault git-pull group/repo -r origin -b main
./blackvault git-pull group/repo --rebase
./blackvault git-push group/repo -r origin -b main
./blackvault git-push group/repo -r origin -b main -u
./blackvault git-push group/repo -r origin -b main --force
./blackvault git-branch-create group/repo new-branch
./blackvault git-branch-switch group/repo main
./blackvault git-branch-rename group/repo new-name
./blackvault git-branch-delete group/repo old-branch
./blackvault git-branch-delete group/repo old-branch --force
./blackvault git-branch-set-upstream group/repo origin/main
./blackvault git-merge group/repo feature/my-feature --no-ff -m "Merge feature"
./blackvault git-remote list group/repo
./blackvault git-remote add group/repo origin https://...
./blackvault git-remote remove group/repo other
./blackvault git-remote set-url group/repo origin https://...
./blackvault git-flow init group/repo
./blackvault git-flow feature-start group/repo my-feature
./blackvault git-flow release-start group/repo 1.2.0
./blackvault git-flow release-finish group/repo 1.2.0 --tag
./blackvault git-flow hotfix-start group/repo 1.2.1
./blackvault git-flow hotfix-finish group/repo 1.2.1 --tag

# Config & serve
./blackvault config get
./blackvault config set git_path /path/to/git
./blackvault install-git
./blackvault serve --port 50051
```

- **Git setting:** ถ้าไม่ตั้ง `git_path` จะใช้ **git ของระบบ (PATH)** เป็น default ถ้าไม่มีจะดูที่ `~/.blackvault/tools/git` (portable). ตั้งเอง: `blackvault config set git_path <path>` หรือรัน `blackvault install-git` แล้วเอา portable ไปวางแล้วตั้ง path
- **Dynamic port:** ถ้า port ชน จะลอง port ถัดไปอัตโนมัติ และเขียนพอร์ตที่ใช้ลง `~/.blackvault/grpc_port` (ลบเมื่อ serve ปิด) — GUI อ่านไฟล์นี้เพื่อเชื่อมต่อ

Config อยู่ที่ `~/.blackvault/config.yaml` (สร้างโดยอัตโนมัติเมื่อไม่มี หรือใช้จาก lib).
