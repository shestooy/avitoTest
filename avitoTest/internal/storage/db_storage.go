package storage

import (
	f "avitoTest/avitoTest/internal/flags"
	"avitoTest/avitoTest/internal/server/model"
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

var Storage DB

//TODO исправить проверки доступа согласно заданию для методов отвечающих за получение тендеров для пользователя
//TODO поменять в проверках ID пользователя на username
//TODO доабавить кастомные ошибки

type DB struct {
	dbPool *pgxpool.Pool
}

func (p *DB) InitDB(ctx context.Context) error {
	var err error
	p.dbPool, err = pgxpool.New(ctx, f.GetPostgresConn())
	if err != nil {
		return err
	}
	return nil
}

func (p *DB) Ping(ctx context.Context) error {
	return p.dbPool.Ping(ctx)
}

func (p *DB) AddNewTender(ctx context.Context, tender model.Tender) (model.Tender, error) {
	var userId string
	err := p.dbPool.QueryRow(ctx, `SELECT user_id FROM organization_responsible
               WHERE organization_id = $1 AND user_id = $2`, tender.OrganizationID, tender.UsernameID).Scan(&userId)

	if err != nil {
		return model.Tender{}, err
	}
	err = p.dbPool.QueryRow(ctx, `INSERT INTO tenders (id, name, organization_id, description, status, servicetype, version, createdat)
			VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7)
			RETURNING id, version, createdat`,
		tender.Name, tender.OrganizationID, tender.Description, tender.Status, tender.ServiceType, 1, time.Now()).
		Scan(&tender.ID, &tender.Version, &tender.CreatedAt)
	return tender, err
}

func (p *DB) UpdateTenderStatus(ctx context.Context, tenderID, userId, status string) (model.Tender, error) {
	var validUserId string
	err := p.dbPool.QueryRow(ctx, `SELECT user_id FROM organization_responsible o
            JOIN tenders t ON t.organization_id = o.organization_id
            WHERE t.id = $1 AND o.user_id = $2`, tenderID, userId).Scan(&validUserId)

	if err != nil {
		return model.Tender{}, err
	}
	var t model.Tender
	err = p.dbPool.QueryRow(context.Background(), `UPDATE tenders SET status = $1 WHERE id = $2 
            RETURNING id, name, organization_id, description, status, servicetype, version, createdat`, status, tenderID).
		Scan(&t.ID, &t.Name, &t.OrganizationID, &t.Description, &t.Status, &t.ServiceType, &t.Version, &t.CreatedAt)
	return t, err
}

func (p *DB) UpdateTender(ctx context.Context, tender model.Tender) (model.Tender, error) {
	var userId string
	err := p.dbPool.QueryRow(ctx, `SELECT user_id FROM organization_responsible
               WHERE organization_id = $1 AND user_id = $2`, tender.OrganizationID, tender.UsernameID).Scan(&userId)

	if err != nil {
		return model.Tender{}, err
	}
	var oldTender model.Tender
	err = p.dbPool.QueryRow(ctx, `SELECT id, name, organization_id, description, status, servicetype, version, createdat FROM tenders
				WHERE id = $1 ORDER BY version DESC LIMIT 1`, tender.Version).Scan(&oldTender)

	if err != nil {
		return model.Tender{}, err
	}

	compare := func(old, new string) string {
		if new != "" {
			return new
		}
		return old
	}
	tender.Name = compare(oldTender.Name, tender.Name)
	tender.Status = compare(oldTender.Status, tender.Status)
	tender.Description = compare(oldTender.Description, tender.Description)
	tender.ServiceType = compare(oldTender.ServiceType, tender.ServiceType)
	tender.Version++

	err = p.dbPool.QueryRow(ctx, `INSERT INTO tenders (id, name, organization_id, description, status, servicetype, version, createdat) 
			VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7)
			RETURNING id, name, organization_id, description, status, servicetype, version, createdat`,
		tender.Name, tender.OrganizationID, tender.Description, tender.Status,
		tender.ServiceType, tender.Version, time.Now()).Scan(&tender.ID, &tender.Name, &tender.OrganizationID,
		&tender.Description, &tender.Status, &tender.ServiceType, &tender.Version, &tender.CreatedAt)
	return tender, err
}

func (p *DB) RollbackTender(ctx context.Context, tenderID, version, username string) (model.Tender, error) {
	var validUserId string
	err := p.dbPool.QueryRow(ctx, `SELECT user_id FROM organization_responsible o
            JOIN tenders t ON t.organization_id = o.organization_id
            WHERE t.id = $1 AND o.user_id = $2`, tenderID, username).Scan(&validUserId)

	if err != nil {
		return model.Tender{}, err
	}
	if version == "1" {
		return model.Tender{}, errors.New("SOON")
	}
	//TODO сюда select запрос к субд
	_, err = p.dbPool.Exec(ctx, `DELETE FROM tenders WHERE id = $1 AND version > $2`, tenderID, version)
	if err != nil {
		return model.Tender{}, err
	}
	//TODO сюда insert запрос к субд
	return model.Tender{}, err
}
