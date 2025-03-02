package dbrepo 

import (
	"coffee-svc/internal/models"
	"context"
	"log"
	"strconv"
)

func (m *PostgresDBRepo) AllProcesses() ([]*models.Process, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select *
		from processes
		order by name
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("error no able to query database for processes")
		return nil, err
	}

	defer rows.Close()

	var processes []*models.Process

	for rows.Next() {
		var process models.Process
		err := rows.Scan(
			&process.ProcessId,
			&process.Name,
			&process.Description,
			&process.TimeDuration,
		)
		if err != nil {
			printRows(rows)
			log.Panic(err)
			return nil, err
		}

		processes = append(processes, &process)
	}

	return processes, nil
}

func (m *PostgresDBRepo) Process(id string) ([]*models.Process, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		select *  
		from processes 
		where processid = $1
		order by name
		limit 1 
	`

	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		log.Println("error no able to query databaases for process")
		return nil, err
	}

	defer rows.Close()

	var processes []*models.Process

	for rows.Next() {
		var process models.Process
		err := rows.Scan(
			&process.ProcessId,
			&process.Name,
			&process.Description,
			&process.TimeDuration,
		)
		if err != nil {
			printRows(rows)
			log.Panic(err)
			return nil, err
		}

		processes = append(processes, &process)
	}
	return processes, nil	
}

func (m *PostgresDBRepo) PostProcess(process *models.Process) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	id := 0
	query := `
	INSERT INTO public.processes (Name, Description, TimeDuration ) VALUES
	($1, $2, $3)
	RETURNING processid`
	err := m.DB.QueryRowContext(ctx, query,
		process.Name,
		process.Description,
		process.TimeDuration).Scan(&id)

	if err != nil {
		log.Println("error no able to query databaases for processes")
		log.Panic(err)
		return "", err
	}

	processid := "New process created with id:" + strconv.FormatInt(int64(id), 10)
	return processid, nil
}
