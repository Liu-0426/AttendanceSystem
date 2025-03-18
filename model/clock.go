package model

import (
    "database/sql"
    "time"
)

type Attendance struct {
    ID       int
    UserID   int
    ClockIn  sql.NullTime
    ClockOut sql.NullTime
    Date     time.Time
}

func SaveClockIn(userID int) error {
    _, err := db.Exec("INSERT INTO attendance (user_id, clock_in, date) VALUES (?, ?, ?)", userID, time.Now(), time.Now().Format("2006-01-02"))
    if err != nil {
        return err
    }
    return nil
}
func SaveClockOut(userID int) error {
    _, err := db.Exec("UPDATE attendance SET clock_out = ? WHERE user_id = ? AND date = ?", time.Now(), userID, time.Now().Format("2006-01-02"))
    if err != nil {
        return err
    }
    return nil
}

func HasClockedIn(userID int) (bool, error) {
    var count int
    err := db.QueryRow("SELECT COUNT(*) FROM attendance WHERE user_id = ? AND date = ?", userID, time.Now().Format("2006-01-02")).Scan(&count)
    if err != nil {
        return false, err
    }
    return count > 0, nil
}

func HasClockedOut(userID int) (bool, error) {
    var count int
    err := db.QueryRow("SELECT COUNT(*) FROM attendance WHERE user_id = ? AND date = ? AND clock_out IS NOT NULL", userID, time.Now().Format("2006-01-02")).Scan(&count)
    if err != nil {
        return false, err
    }
    return count > 0, nil
}

func GetClockInsByUserID(userID int) ([]Attendance, error) {
    var attendances []Attendance
    rows, err := db.Query("SELECT id, user_id, clock_in, clock_out, date FROM attendance WHERE user_id = ?", userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
        var attendance Attendance
        err := rows.Scan(&attendance.ID, &attendance.UserID, &attendance.ClockIn, &attendance.ClockOut, &attendance.Date)
        if err != nil {
            return nil, err
        }
        attendances = append(attendances, attendance)
    }
    return attendances, nil
}

func GetAllClockList() ([]Attendance, error) {
    var attendances []Attendance
    rows, err := db.Query("SELECT id, user_id, clock_in, clock_out, date FROM attendance")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
        var attendance Attendance
        err := rows.Scan(&attendance.ID, &attendance.UserID, &attendance.ClockIn, &attendance.ClockOut, &attendance.Date)
        if err != nil {
            return nil, err
        }
        attendances = append(attendances, attendance)
    }
    return attendances, nil
}
func GetTodayClockInCount() (int, error) {
    var count int
    err := db.QueryRow("SELECT COUNT(*) FROM attendance WHERE date = ?", time.Now().Format("2006-01-02")).Scan(&count)
    if err != nil {
        return 0, err
    }
    return count, nil
}