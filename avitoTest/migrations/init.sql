CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS employee (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(255) NOT NULL,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);


CREATE INDEX idx_employee_username ON employee (username);

CREATE TYPE organization_type AS ENUM (
    'IE',
    'LLC',
    'JSC'
    );


CREATE TABLE organization (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    type organization_type,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tenders (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    organization_id UUID NOT NULL,
    description TEXT,
    status VARCHAR(50) NOT NULL,
    serviceType VARCHAR(50),
    version INT NOT NULL,
    createdAt TIMESTAMP NOT NULL
);

CREATE INDEX idx_tenders_organization_id ON tenders (organization_id);

ALTER TABLE tenders
    ADD CONSTRAINT fk_tenders_organization FOREIGN KEY (organization_id) REFERENCES organization (id);

CREATE TABLE organization_responsible (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID NOT NULL,
    user_id UUID NOT NULL
);

CREATE INDEX idx_org_responsible_org_id ON organization_responsible (organization_id);
CREATE INDEX idx_org_responsible_user_id ON organization_responsible (user_id);

ALTER TABLE organization_responsible
    ADD CONSTRAINT fk_org_responsible_org FOREIGN KEY (organization_id) REFERENCES organization (id),
    ADD CONSTRAINT fk_org_responsible_user FOREIGN KEY (user_id) REFERENCES employee (id);


CREATE TABLE IF NOT EXISTS offers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL,
    authorType VARCHAR(50) NOT NULL,
    author_Id UUID NOT NULL,
    version INT NOT NULL,
    createdAt TIMESTAMP NOT NULL,
    creatorUsername VARCHAR(255),
    organization_id UUID NOT NULL
);

CREATE INDEX idx_offers_organization_id ON offers (organization_id);
CREATE INDEX idx_offers_author_id ON offers (author_Id);
CREATE INDEX idx_offers_creator_username ON offers (creatorUsername);

ALTER TABLE offers
    ADD CONSTRAINT fk_offers_author FOREIGN KEY (author_Id) REFERENCES employee (id),
    ADD CONSTRAINT fk_offers_organization FOREIGN KEY (organization_id) REFERENCES organization (id);


CREATE TABLE IF NOT EXISTS approvals (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    rej BOOLEAN NOT NULL,
    approve BOOLEAN NOT NULL,
    offers_id UUID NOT NULL,
    decision_maker_id UUID NOT NULL
);

CREATE INDEX idx_approvals_offers_id ON approvals (offers_id);
CREATE INDEX idx_approvals_decision_maker_id ON approvals (decision_maker_id);


ALTER TABLE approvals
    ADD CONSTRAINT fk_approvals_offers FOREIGN KEY (offers_id) REFERENCES offers (id),
    ADD CONSTRAINT fk_approvals_decision_maker FOREIGN KEY (decision_maker_id) REFERENCES employee (id);


CREATE TABLE IF NOT EXISTS feedback (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL,
    authorType VARCHAR(50) NOT NULL,
    authorId UUID NOT NULL,
    version INT NOT NULL,
    createdAt TIMESTAMP NOT NULL,
    bidId UUID NOT NULL
);

CREATE INDEX idx_feedback_author_id ON feedback (authorId);
CREATE INDEX idx_feedback_bid_id ON feedback (bidId);


ALTER TABLE  feedback
    ADD CONSTRAINT fk_feedback_author FOREIGN KEY (authorId) REFERENCES employee (id),
    ADD CONSTRAINT fk_feedback_bid FOREIGN KEY (bidId) REFERENCES tenders (id);



