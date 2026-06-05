| Skill                       | Công dụng       |
| --------------------------- | --------------- |
| understand                  | quét & phân tích toàn bộ project |
| understand-dashboard        | mở UI biểu đồ trực quan |
| understand-chat             | hỏi đáp về cấu trúc code |
| understand-explain          | giải thích 1 file/hàm cụ thể |
| understand-diff             | kiểm tra ảnh hưởng trước khi commit |
| understand-onboard          | tạo tài liệu cho dev mới |
| understand-domain           | phân tích business logic (nghiệp vụ) |
___
# Understand Anything Guide (Tiếng Việt)

## Giới thiệu

Đây là bộ công cụ giúp bạn hiểu codebase siêu to khổng lồ mà không cần đọc từng dòng code.

Mục tiêu:
* nhìn được bức tranh tổng thể (Big Picture)
* không code mò mẫm
* biết file này liên quan đến file nào
* onboarding dev mới nhanh gọn

---

# Workflow chuẩn

```text
Use /understand (chỉ 1 lần đầu hoặc khi đổi code nhiều)
↓
Use /understand-dashboard (để xem kiến trúc)
↓
Use /understand-chat (hỏi đáp chung) 
hoặc /understand-explain (hỏi 1 file cụ thể)
↓
Code thay đổi
↓
Use /understand-diff (check impact)
```

---

# 1. understand

## Dùng khi nào?
* mới join project
* lần đầu tiên chạy tool
* cấu trúc thư mục vừa bị thay đổi lớn
* muốn ép AI cập nhật lại kiến thức về codebase

---

## Template chung
```text
/understand

(Nếu muốn tiếng Việt trôi chảy)
/understand --language vi
```

---

# 2. understand-dashboard

## Dùng khi nào?
* đọc code mãi không hiểu flow
* muốn xem diagram trực quan
* cần thuyết trình kiến trúc cho sếp/team

---

## Template chung
```text
/understand-dashboard
```

---

# 3. understand-chat

## Dùng khi nào?
* tìm file xử lý logic
* không biết bắt đầu sửa từ đâu
* trace bug liên quan nhiều module

---

## Template chung
```text
/understand-chat <câu hỏi của bạn>
```

## Ví dụ thực tế
```text
/understand-chat Flow đăng nhập hiện tại đang đi qua những file nào?
/understand-chat Chỗ nào đang handle việc upload hình ảnh?
```

---

# 4. understand-explain

## Dùng khi nào?
* file quá dài (VD: App.tsx 70KB)
* hàm quá phức tạp
* đọc code cũ của người khác (legacy code)

---

## Template chung
```text
/understand-explain <đường dẫn file>
```

## Ví dụ thực tế
```text
/understand-explain src/App.tsx
/understand-explain server/internal/handlers/prompt_builder.go
```

---

# 5. understand-diff

## Dùng khi nào?
* chuẩn bị commit/tạo Pull Request
* vừa sửa 1 hàm dùng chung (utils/helpers)
* sợ sửa chỗ này chết chỗ kia 💀

---

## Template chung
```text
/understand-diff
```

## Ví dụ thực tế
```text
/understand-diff
Phân tích xem những thay đổi tôi vừa làm ở auth middleware có làm sập các API khác không?
```

---

# 6. understand-onboard

## Dùng khi nào?
* có dev mới vào team
* bàn giao project
* viết docs nhanh

---

## Template chung
```text
/understand-onboard
Tạo một file hướng dẫn setup và luồng chạy chính cho người mới.
```

---

# Kết luận

Đừng đọc code như đọc truyện (đọc từ trên xuống dưới).
Hãy đọc code theo Graph (hiểu từ bức tranh tổng thể xuống chi tiết).

* Chưa hiểu tổng thể → `/understand-dashboard`
* Chưa hiểu flow → `/understand-chat`
* Chưa hiểu 1 file → `/understand-explain`
* Sợ sửa lỗi → `/understand-diff`
---
GitHub của plugin có bản cập nhật mới
Bước 1: Tải mã nguồn mới nhất từ GitHub
cd C:\Users\anhnq.568E-local\.understand-anything\repo\understand-anything-plugin
git pull

Bước 2: Cài đặt lại thư viện và Build lại bộ lõi (Quan trọng) Do mã nguồn có thể thay đổi các gói phụ thuộc hoặc mã lõi, bạn nên chạy lại lệnh cài đặt:
pnpm install
pnpm --filter @understand-anything/core build

Bước 3: Cập nhật vào dự án của bạn (Tùy trường hợp)

Trường hợp 1 (Chỉ cập nhật code của các skill cũ): Nếu repo chỉ cập nhật mã nguồn bên trong các skill đã có sẵn, bạn không cần làm gì thêm. Các liên kết (Junctions) bạn đã tạo sẽ tự động phản ánh các thay đổi này ngay lập tức.
Trường hợp 2 (Repo có thêm các skill hoàn toàn mới): Nếu tác giả thêm một thư mục skill mới, bạn chỉ cần chạy lại câu lệnh PowerShell mà bạn vừa hỏi ở trên (hoặc copy thư mục skill mới đó) để nó quét và tạo liên kết Junction bổ sung cho skill mới vào trong thư mục D:\Code\AI\playable.ai\.agent\skills\ của bạn.