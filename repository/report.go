package repository

import (
	"Book_Cart_Project/models"
	"log"
)

func (r Repository) AdminReport(mnth models.ReportIn) ([]models.OrderReport, error) {

	var report []models.OrderReport
	query := `SELECT 
	DATE_TRUNC('month',order_at) AS order_in_month, total_price,
	COUNT(order_id) AS order_count
	FROM user_order
	WHERE order_status = $1
	GROUP BY DATE_TRUNC('month',order_at), total_price;`

	rows, err := r.DB.Query(query, mnth.Order_status)
	if err != nil {
		log.Println("report failes")
		return report, err
	}
	defer rows.Close()

	for rows.Next() {
		var reports models.OrderReport
		err := rows.Scan(
			&reports.Order_month,
			&reports.Total_price,
			&reports.Order_count,
		)
		if err != nil {
			log.Println("report scan failed", err)
			return report, err
		}
		report = append(report, reports)
		if err = rows.Err(); err != nil {
			return report, err
		}
	}
	return report, nil
}

func (r Repository) EditOrderStatus(status models.ChangeOrder) (models.ChangeOrder, error) {
	query := `UPDATE 
					user_order
				SET 
					order_status = $1
				WHERE 
					order_id = $2
				RETURNING 
					order_id, 
					order_status;`
	err := r.DB.QueryRow(
		query,
		status.Order_status,
		status.Order_id,
	).Scan(
		&status.Order_id,
		&status.Order_status,
	)
	if err != nil {
		return status, err
	}
	return status, nil
}
