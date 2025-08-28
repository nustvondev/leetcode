# 📅 Daily Practice Log

Nhật ký luyện tập LeetCode & DSA.  
Mỗi ngày mình sẽ ghi lại: **Ngày / Bài tập / Ý tưởng / Ghi chú**.  
👉 Mục tiêu: duy trì thói quen giải 1-2 bài/ngày, ưu tiên hiểu **tư duy & trade-off** giữa các giải pháp.

---

## Week 1

### Day 01 (2025-08-26)
- **Top75**: Two Sum ✅
  - **Ý tưởng**:
    - **Brute Force**:  
      - Duyệt qua tất cả các cặp số `(i, j)`.  
      - Nếu `nums[i] + nums[j] == target` → return kết quả.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n^2)`  
        - 💾 Không gian: `O(1)`  
    - **Hash Map (Tối ưu)**:  
      - Duyệt qua từng số, tính `complement = target - num`.  
      - Kiểm tra nếu `complement` đã có trong hash map → return chỉ mục.  
      - Nếu chưa có → thêm `num` vào hash map với chỉ mục của nó.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n)`  
        - 💾 Không gian: `O(n)`  
  - **Ghi chú**:  
    - Hash Map là giải pháp tối ưu.  
    - Bài toán này luyện khả năng dùng hash map để kiểm tra **sự tồn tại** và **tìm kiếm nhanh**.

---

### Day 02 (2025-08-27)
- **Top75**: Contains Duplicate ✅
  - **Ý tưởng**:
    - **Brute Force**:  
      - So sánh từng cặp phần tử `(i, j)`.  
      - Nếu có `nums[i] == nums[j]` → return `true`.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n^2)`  
        - 💾 Không gian: `O(1)`  
    - **Hash Map / Hash Set (Tối ưu)**:  
      - Dùng hash map/set để lưu các số đã gặp.  
      - Nếu gặp số đã tồn tại trong map/set → return `true`.  
      - Nếu duyệt hết không trùng → return `false`.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n)`  
        - 💾 Không gian: `O(n)`  
    - **Sorting**:  
      - Sắp xếp mảng trước.  
      - Duyệt kiểm tra xem có phần tử nào giống phần tử trước đó.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n log n)`  
        - 💾 Không gian: phụ thuộc vào thuật toán sort (trung bình `O(1)` đến `O(n)`).  
  - **Ghi chú**:  
    - Hash Set là giải pháp tối ưu nhất cho thời gian.  
    - Sorting hữu ích khi cần tiết kiệm bộ nhớ hoặc muốn xử lý theo thứ tự.  
    - Brute Force chỉ để tham khảo, không nên dùng thực tế.  

---

### Day 03 (2025-08-28)
- **Top75**: Best Time to Buy and Sell Stock ✅
  - **Ý tưởng**:
    - **Brute Force**:  
      - Duyệt qua mọi cặp `(buy, sell)` với `buy < sell`.  
      - Tính `profit = prices[sell] - prices[buy]`.  
      - Giữ lại lợi nhuận lớn nhất.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n^2)`  
        - 💾 Không gian: `O(1)`  
    - **One-Pass (Tối ưu)**:  
      - Duyệt mảng 1 lần, theo dõi **minPrice** (giá nhỏ nhất từ đầu đến hiện tại).  
      - Với mỗi giá `price`, tính `profit = price - minPrice`.  
      - Cập nhật `maxProfit` nếu `profit` lớn hơn.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n)`  
        - 💾 Không gian: `O(1)`  
  - **Ghi chú**:  
    - Đây là bài “khởi đầu” cho chuỗi bài **Stock Problems**.  
    - Tư duy quan trọng: **theo dõi min giá trước đó** và **so sánh chênh lệch**.  
    - Các biến thể nâng cao (Stock II, III, IV, k transactions) sẽ cần DP.

---
