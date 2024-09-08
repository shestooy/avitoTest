package storage

import (
	f "avitoTest/avitoTest/internal/flags"
	"avitoTest/avitoTest/internal/server/model"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

var Storage DB

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
