# Diagram

- Task Steps
```mermaid
flowchart LR
    s1[Step 1]
    s21[Step 2 sub 1]
    s22[Step 2 sub 2]
    s23[Step 2 sub 3]
    s24[Step 2 sub 4]
    s25[Step 2 sub 5]
    s3[Step 3]
    s4[Step 4]

    s1 -->|fan-out| s21 -->|fan-in| s3
    s1 -->|fan-out| s22 -->|fan-in| s3
    s1 -->|fan-out| s23 -->|fan-in| s3
    s1 -->|fan-out| s24 -->|fan-in| s3
    s1 -->|fan-out| s25 -->|fan-in| s3

    s3 --> s4
```