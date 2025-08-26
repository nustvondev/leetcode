# 📅 Daily Practice Log

Nhật ký luyện tập LeetCode & DSA.  
Mỗi ngày mình sẽ ghi lại: **Ngày / Bài tập / Ý tưởng / Ghi chú**.

---

## Week 1

### Day 01 (2025-08-26)
- **Top75**: Two Sum ✅
  - **Ý tưởng**:
    - Brute Force: Duyệt qua tất cả các cặp số. Độ phức tạp thời gian O(n^2), không gian O(1).
    - Hash Map: Sử dụng hash map để lưu trữ số và chỉ mục của nó. Tìm `complement` (target - số hiện tại) trong hash map. Độ phức tạp thời gian O(n), không gian O(n).
  - **Ghi chú**: Hash Map là giải pháp tối ưu hơn về thời gian.

---
