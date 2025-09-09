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

### Day 04 (2025-08-29)
- **Top75**: Valid Anagram ✅
  - **Ý tưởng**:
    - **Brute Force (Sort + So sánh)**:  
      - Sắp xếp 2 chuỗi rồi so sánh từng ký tự.  
      - Nếu giống nhau → là anagram.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n log n)`  
        - 💾 Không gian: `O(1)` hoặc `O(n)` tùy sort.  
    - **Hash Map (Tối ưu cho lowercase)**:  
      - Đếm tần suất ký tự trong chuỗi `s`.  
      - Trừ đi tần suất ký tự trong chuỗi `t`.  
      - Nếu có ký tự nào âm → không phải anagram.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n)`  
        - 💾 Không gian: `O(1)` (26 ký tự tiếng Anh).  
    - **Unicode-Friendly (Tổng quát)**:  
      - Sử dụng hash map để đếm tần suất ký tự dạng `rune`.  
      - Áp dụng cho mọi ngôn ngữ, kể cả Unicode.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n)`  
        - 💾 Không gian: `O(k)` với `k` là số ký tự khác nhau.  
  - **Ghi chú**:  
    - Hash Map là cách tối ưu cho input chữ thường.  
    - Với Unicode, cần dùng map rune để xử lý tổng quát.  
    - Bài này rèn luyện kỹ năng **count frequency** – nền tảng cho nhiều bài liên quan đến chuỗi.  

---


### Day 05 (2025-09-03)
- **Top75**: Valid Parentheses ✅
  - **Ý tưởng**:
    - **Brute Force (String Replace)**:  
      - Liên tục thay thế các cặp `"()"`, `"{}"`, `"[]"` bằng chuỗi rỗng.  
      - Nếu cuối cùng chuỗi rỗng → hợp lệ.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n^2)` trong trường hợp xấu.  
        - 💾 Không gian: `O(n)` do tạo chuỗi mới.  
    - **Stack (Chuẩn – Best Practice)**:  
      - Duyệt từng ký tự trong chuỗi:  
        - Nếu là ngoặc mở → push vào stack.  
        - Nếu là ngoặc đóng → kiểm tra top của stack có khớp không.  
      - Cuối cùng stack rỗng → hợp lệ.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n)`  
        - 💾 Không gian: `O(n)`  
    - **Optimized Stack (Push ngoặc cần đóng)**:  
      - Khi gặp `(` thì push `)`, gặp `[` thì push `]`, gặp `{` thì push `}`.  
      - Khi gặp ngoặc đóng, chỉ cần so với top stack.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n)`  
        - 💾 Không gian: `O(n)`  
  - **Ghi chú**:  
    - Stack là giải pháp tối ưu và dễ cài đặt nhất.  
    - Cách push “ngoặc cần đóng” giúp code gọn hơn.  
    - Đây là bài rèn luyện tư duy **stack ứng dụng cho kiểm tra cấu trúc** (parsing, compiler, expression).  


---


### Day 06 (2025-09-04)
- **Top75**: Maximum Subarray ✅
  - **Ý tưởng**:
    - **Brute Force (O(n³))**:  
      - Duyệt tất cả các subarray `(i..j)`.  
      - Tính tổng từng subarray, lấy max.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n^3)`  
        - 💾 Không gian: `O(1)`  
    - **Prefix Sum (O(n²))**:  
      - Tính prefix sum: `prefix[i] = tổng nums[0..i-1]`.  
      - Subarray sum `(i..j) = prefix[j+1] - prefix[i]`.  
      - Giảm 1 vòng lặp so với brute force.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n^2)`  
        - 💾 Không gian: `O(n)`  
    - **Kadane’s Algorithm (O(n)) – Best Practice**:  
      - Duyệt mảng 1 lần:  
        - Với mỗi `nums[i]`, chọn giữa bắt đầu subarray mới (`nums[i]`) hoặc nối tiếp subarray trước (`currSum + nums[i]`).  
        - Cập nhật `maxSum`.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n)`  
        - 💾 Không gian: `O(1)`  
    - **Divide & Conquer (O(n log n))**:  
      - Chia đôi mảng thành left và right.  
      - Kết quả tối đa là max của:  
        1. Subarray max bên trái  
        2. Subarray max bên phải  
        3. Subarray max đi qua mid  
      - Dùng đệ quy để giải từng phần.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n log n)`  
        - 💾 Không gian: `O(log n)` (stack đệ quy)  
  - **Ghi chú**:  
    - Kadane’s là chuẩn nhất (O(n), O(1)).  
    - Brute force & prefix sum giúp hiểu bản chất nhưng không thực tế.  
    - Divide & Conquer hữu ích trong phỏng vấn để thể hiện tư duy thuật toán.  
    - Đây là bài cực kỳ kinh điển, thường được hỏi trong interview để kiểm tra khả năng tối ưu tư duy.  


---


### Day 07 (2025-09-05)
- **Top75**: Product of Array Except Self ✅
  - **Ý tưởng**:
    - **Brute Force (O(n²))**:  
      - Với mỗi `i`, tính tích tất cả phần tử trừ `nums[i]`.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n²)`  
        - 💾 Không gian: `O(1)`  
    - **Prefix & Suffix Arrays (O(n), O(n))**:  
      - Lưu `prefix[i] = tích từ trái`.  
      - Lưu `suffix[i] = tích từ phải`.  
      - Kết quả = `prefix[i] * suffix[i]`.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n)`  
        - 💾 Không gian: `O(n)`  
    - **Optimized Prefix-Suffix (O(n), O(1)) – Best Practice**:  
      - Dùng `res[]` để lưu prefix.  
      - Một vòng từ phải sang trái để nhân suffix.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n)`  
        - 💾 Không gian: `O(1)` (không tính `res`).  
  - **Ghi chú**:  
    - Đây là dạng bài kinh điển cho kỹ thuật **prefix-suffix**.  
    - Hạn chế “không dùng chia” giúp luyện tư duy tối ưu.  
    - Phiên bản O(1) space là chuẩn để nhớ và dùng trong interview.  

---


### Day 08 (2025-09-08)
- **Top75**: 3Sum ✅
  - **Ý tưởng**:
    - **Brute Force (O(n³))**:  
      - Duyệt tất cả bộ ba `(i, j, k)`.  
      - Nếu tổng = 0 thì lưu kết quả (dùng set để tránh trùng).  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n³)`  
        - 💾 Không gian: `O(n²)` để lưu kết quả.  
    - **Sorting + Two Pointers (O(n²)) – Best Practice**:  
      - Sort mảng trước.  
      - Với mỗi `nums[i]`, dùng hai con trỏ `l` và `r` để tìm `nums[l] + nums[r] = -nums[i]`.  
      - Nếu sum = 0 → lưu triplet, di chuyển con trỏ và skip duplicates.  
      - Nếu sum < 0 → tăng `l`.  
      - Nếu sum > 0 → giảm `r`.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n²)`  
        - 💾 Không gian: `O(1)` (không tính output).  
  - **Ghi chú**:  
    - Đây là bài kinh điển dùng **Two Pointers + Sorting**.  
    - Cần cẩn thận xử lý **duplicate triplets**.  
    - Là nền tảng cho các biến thể nâng cao: **4Sum**, **k-Sum**, …  

---


---

## 3. Daily Note (📅 Day 09 – 2025-09-09)

```markdown
### Day 09 (2025-09-09)
- **Top75**: Merge Intervals ✅
  - **Ý tưởng**:
    - **Brute Force (O(n²))**:  
      - So sánh từng cặp interval để xem có overlap không.  
      - Nếu có → merge, lặp lại cho đến khi không merge được nữa.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n²)`  
        - 💾 Không gian: `O(n)`  
    - **Sorting + Greedy (O(n log n)) – Best Practice**:  
      - Sort intervals theo `start`.  
      - Duyệt tuần tự, so với interval cuối trong `res`:  
        - Nếu overlap → merge lại.  
        - Nếu không → push interval mới vào.  
      - **Độ phức tạp**:  
        - ⏱️ Thời gian: `O(n log n)` (do sort).  
        - 💾 Không gian: `O(1)` (nếu tái sử dụng input).  
  - **Ghi chú**:  
    - Đây là bài cơ bản về **interval problems**.  
    - Pattern thường gặp: **sort + merge**.  
    - Quan trọng cho các bài nâng cao: **Insert Interval, Meeting Rooms, Employee Free Time**.  
