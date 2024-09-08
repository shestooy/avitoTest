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

	_, err = p.dbPool.Exec(ctx, `INSERT INTO tenders (id, name, organization_id, description, status, servicetype, version, createdat)
				VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7)`,
		tender.Name, tender.OrganizationID, tender.Description, tender.Status, tender.ServiceType, 1, time.Now())
	return model.Tender{}, err
}

func (p *DB) UpdateTenderStatus(ctx context.Context, tenderID, userId, status string) error {
	var validUserId string
	err := p.dbPool.QueryRow(ctx, `SELECT user_id FROM organization_responsible o
            JOIN tenders t ON t.organization_id = o.organization_id
            WHERE t.id = $1 AND o.user_id = $2`, tenderID, userId).Scan(&validUserId)

	if err != nil {
		return err
	}
	_, err = p.dbPool.Exec(context.Background(), `
            UPDATE tenders
            SET status = $1
            WHERE id = $2
        `, status, tenderID)
	return err
}
