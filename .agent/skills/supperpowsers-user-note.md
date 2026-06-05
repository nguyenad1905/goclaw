| Skill                       | Công dụng       |
| --------------------------- | --------------- |
| brainstorming               | phân tích trước |
| writing-plans               | chia task       |
| executing-plans             | code từng bước  |
| systematic-debugging        | debug           |
| test-driven-development     | TDD             |
| requesting-code-review      | review code     |
| using-git-worktrees         | branch/worktree |
| subagent-driven-development | multi-agent     |
___
# Superpowers Prompt Templates Guide (Tiếng Việt)

## Giới thiệu

Đây là bộ template prompt:

Mục tiêu:

* ép AI làm việc có quy trình
* tránh code hấp tấp
* giảm hallucination
* giảm bug
* giảm phá project 😏

---

# Workflow chuẩn

```text
Use brainstorming
↓
Use writing-plans
↓
Approve plan
↓
Use executing-plans
↓
Use test-driven-development
↓
Use requesting-code-review
```

---

# 1. brainstorming

## Dùng khi nào?

* bắt đầu project mới
* thiết kế architecture
* feature phức tạp
* chưa rõ requirement
* refactor lớn

---

## Template chung

```text
Use brainstorming first.

Phân tích yêu cầu trước.
Không code ngay.

Hãy:
- phân tích architecture
- xác định constraint
- xác định edge case
- xác định risk
- đề xuất nhiều approach
- giải thích tradeoff

Chờ tôi approve trước khi implement.
```

---

## Ví dụ thực tế — auth system

```text
Use brainstorming first.

Phân tích hệ thống authentication và authorization hiện tại.

Mục tiêu:
- kiểm tra flow auth
- kiểm tra permission
- tìm security risk
- đánh giá architecture

Không code ngay.

Hãy:
- phân tích flow hiện tại
- xác định điểm yếu
- tìm IDOR risk
- tìm privilege escalation risk
- đề xuất cách cải thiện
```

---

# 2. writing-plans

## Dùng khi nào?

* trước khi implement feature
* trước khi refactor
* thay đổi lớn
* thay đổi database
* thay đổi auth

---

## Template chung

```text
Use writing-plans.

Trước khi code:
- phân tích codebase hiện tại
- xác định architecture liên quan
- xác định file sẽ bị ảnh hưởng
- xác định risk và edge case

Sau đó tạo implementation plan chi tiết.

Yêu cầu plan phải có:
1. Mục tiêu
2. Phân tích hiện trạng
3. Các vấn đề/risk
4. Database/schema changes
5. Backend changes
6. Frontend changes
7. API changes
8. Authorization/security checks
9. Step-by-step implementation
10. Test plan
11. Rollback strategy

Không implement ngay.
Chờ tôi approve plan trước.
```

---

## Ví dụ thực tế — project ownership

```text
Use writing-plans.

Phân tích hệ thống auth và authorization hiện tại.

Mục tiêu:
- thêm project-level authorization
- mỗi user chỉ được quản lý project của chính họ

Yêu cầu:
- không code ngay
- phân tích toàn bộ flow auth hiện tại
- xác định middleware/guard đang dùng
- kiểm tra API nào đang thiếu authorization
- xác định nguy cơ IDOR hoặc data leakage

Tạo implementation plan chi tiết bao gồm:
1. Current architecture
2. Current auth flow
3. Current authorization flow
4. Problems found
5. Database/schema changes
6. Backend API changes
7. Frontend permission changes
8. Ownership enforcement strategy
9. Step-by-step implementation
10. Test cases
11. Migration strategy
12. Rollback plan

Mỗi step phải:
- nhỏ
- verify được
- ít risk
```

---

# 3. executing-plans

## Dùng khi nào?

* bắt đầu code thật
* implement theo plan
* rollout feature

---

## Template chung

```text
Use executing-plans.

Implement theo implementation plan đã approve.

Yêu cầu:
- chỉ làm từng step nhỏ
- verify sau mỗi step
- không refactor ngoài scope
- giữ consistency với codebase hiện tại
- tránh thay đổi lớn không cần thiết

Sau mỗi step:
- giải thích thay đổi
- chạy test/check nếu có
- báo risk nếu phát hiện vấn đề
```

---

## Ví dụ thực tế — implement permission

```text
Use executing-plans.

Implement project ownership authorization.

Yêu cầu:
- enforce authorization ở backend
- không chỉ check frontend
- user không được access project của user khác
- protect nested resources
- follow existing architecture
- giữ change nhỏ nhất có thể

Sau mỗi step:
- verify logic
- run tests
- explain changes
```

---

# 4. test-driven-development

## Dùng khi nào?

* business logic
* parser
* API
* pricing engine
* auth
* validation
* algorithm

---

## Template chung

```text
Use test-driven-development.

Yêu cầu:
- viết test fail trước
- implement minimal code để pass test
- refactor sau khi test pass
- không skip test phase

Mỗi step:
1. Red
2. Green
3. Refactor
```

---

## Ví dụ thực tế — permission tests

```text
Use test-driven-development.

Viết authorization tests cho project ownership.

Cần test:
- owner access allowed
- non-owner access denied
- nested resource protection
- API filtering
- update/delete restrictions
- edge cases

Viết test fail trước rồi mới implement.
```

---

# 5. systematic-debugging

## Dùng khi nào?

* bug khó
* random issue
* auth issue
* crawler issue
* flaky tests
* production bug

---

## Template chung

```text
Use systematic-debugging.

Không fix ngẫu nhiên.

Hãy:
1. reproduce bug
2. đọc exact error
3. isolate issue
4. tạo hypothesis
5. verify từng hypothesis
6. implement smallest fix
7. verify fix bằng test

Giải thích:
- root cause
- fix
- risk
```

---

## Ví dụ thực tế — auth bug

```text
Use systematic-debugging.

Có vẻ user đang nhìn thấy dữ liệu của user khác.

Hãy:
- reproduce bug
- kiểm tra API response
- kiểm tra database query
- kiểm tra middleware auth
- kiểm tra ownership validation
- tìm root cause
- đề xuất fix nhỏ nhất
```

---

# 6. requesting-code-review

## Dùng khi nào?

* trước khi merge
* security review
* audit code
* review refactor

---

## Template chung

```text
Use requesting-code-review.

Review code critically.

Tập trung vào:
- bugs
- security risks
- edge cases
- performance issues
- missing tests
- bad practices
- maintainability

Cho tôi:
1. Issues found
2. Severity
3. Suggested fixes
4. Missing tests
5. Merge recommendation
```

---

## Ví dụ thực tế — auth review

```text
Use requesting-code-review.

Review phần authorization vừa implement.

Tập trung vào:
- IDOR vulnerability
- privilege escalation
- missing backend checks
- frontend-only protection
- nested resource access
- ownership bypass
- race conditions
- missing tests
```

---

# 7. subagent-driven-development

## Dùng khi nào?

* project lớn
* monorepo
* auth phức tạp
* nhiều module
* refactor lớn

---

## Template chung

```text
Use subagent-driven-development.

Split work into specialized roles:
- architecture analysis
- security analysis
- backend implementation
- frontend implementation
- testing
- code review

Mỗi role:
- chỉ tập trung nhiệm vụ riêng
- verify công việc của role khác
- report findings trước khi code
```

---

## Ví dụ thực tế — RBAC system

```text
Use subagent-driven-development.

Implement RBAC system.

Tách role:
- architect
- security reviewer
- backend engineer
- frontend engineer
- testing engineer
- final reviewer

Yêu cầu:
- không implement toàn bộ cùng lúc
- verify từng phase
- review chéo giữa các role
```

---

# 8. using-git-worktrees

## Dùng khi nào?

* nhiều feature song song
* tránh phá branch chính
* experiment
* risky refactor

---

## Template chung

```text
Use using-git-worktrees.

Tạo worktree riêng cho feature này.

Yêu cầu:
- không sửa trực tiếp main branch
- isolate changes
- giữ branch sạch
- explain branch strategy
```

---

# Prompt helper cực mạnh 😏

## Ép AI không hấp tấp

```text
Do not code immediately.
```

---

## Ép chia step nhỏ

```text
Each implementation step must:
- be small
- be verifiable
- minimize risk
- avoid large refactors
```

---

## Ép follow architecture hiện tại

```text
Follow the current architecture style.
Avoid unnecessary abstractions.
```

---

## Ép AI verify liên tục

```text
Verify after each step before continuing.
```

---

## Ép backend auth thật sự

```text
Enforce authorization at backend level.
Do not rely only on frontend checks.
```

---

# Prompt anti-thảm-họa 💀

## Đừng dùng kiểu này

```text
build entire system
```

Hoặc:

```text
fix everything
```

AI sẽ:

* rewrite nguyên project
* refactor lung tung
* summon bug thần thánh

---

# Prompt chuẩn senior workflow

```text
Use brainstorming first.
Do not code yet.

Sau khi phân tích:
- tạo implementation plan
- chia step nhỏ
- verify từng step
- viết test trước khi implement logic quan trọng
- review code sau khi hoàn tất
```

---

# Kết luận

AI mạnh nhất khi:

* có constraint rõ
* có workflow rõ
* có verify liên tục
* có review

Không workflow:

* AI = intern tăng động 😏

Có workflow:

* AI bắt đầu giống senior engineer hơn rất nhiều.
