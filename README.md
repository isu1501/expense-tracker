# 📌 Expense Tracker CLI

A simple command-line expense tracker built with Go that stores data in a JSON file.

This application allows users to:

- Add an expense
- Update an expense
- Delete an expense
- View all expenses
- View total expense summary
- View monthly expense summary (current year)

## 🚀 Features

- JSON-based storage (no database required)
- Clean layered architecture (Service + Repository)
- Simple CLI interface
- Monthly filtering support
- Auto-generated incremental IDs

## ▶️ How to Run

Build the application

```bash
go build -o expense-tracker
```

All commands are executed using:

```bash 
./expense-tracker <command>
```

## 📖 Usage

### ➕ Add an Expense
```bash
expense-tracker add --description "Lunch" --amount 20
```

### 📋 List All Expenses
```bash 
expense-tracker list
```

### 📊 View Total Summary
```bash
expense-tracker summary
```

### 📅 View Monthly Summary (Current Year)
```bash
expense-tracker summary --month 8
```

### ❌ Delete an Expense
```bash 
expense-tracker delete --id 2
```

## 💾 Storage

All data is stored locally in expenses.json:

```json
[
  {
    "id": 1,
    "description": "Lunch",
    "amount": 20,
    "created_at": "2024-08-06T12:00:00Z"
  }
]
```

## 🌐 Project Page

Project Repository URL:

```bash 
https://roadmap.sh/projects/expense-tracker
```

## 🔮 Future Improvements

- Update expense
- Export to CSV
- Unit tests