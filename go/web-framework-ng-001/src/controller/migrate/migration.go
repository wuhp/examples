package migrate

import (
    "github.com/BurntSushi/migration"
)

//////////////////////////////////////////////////////////////////////

func Migrate_1(tx migration.LimitedTx) error {
    scripts := []string{
        accountTable,
        repoTable,
        setRepoForeignKey,
        teamTable,
        setTeamForeignKey,
        teammemberTable,
        setTeammemberForeignKeyTeam,
        setTeammemberForeignKeyMember,
        orgmemberTable,
        setOrgmemberForeignKeyOrg,
        setOrgmemberForeignKeyMember,
        teamaccessTable,
        setTeamaccessForeignKeyTeam,
        setTeamaccessForeignKeyRepo,
        tagTable,
        setTagForeignKeyRepo,
        hookTable,
        setHookForeignKeyRepo,
    }   
    
    for _, cmd := range scripts {
        if _, err := tx.Exec(cmd); err != nil {
            return err
        }   
    }   
    
    return nil
} 

// account
var accountTable = `
CREATE TABLE IF NOT EXISTS account (
    id                     INT AUTO_INCREMENT PRIMARY KEY,
    name                   VARCHAR(32),
    display_name           VARCHAR(32),
    is_org                 BOOLEAN,
    create_ts              BIGINT, 
    org_description        MEDIUMBLOB,
    user_type              VARCHAR(16),
    user_password          VARCHAR(64),
    user_email             VARCHAR(64),
    user_global_permission VARCHAR(16),
    UNIQUE (name)
)AUTO_INCREMENT=10000;
`

// repo
var repoTable = `
CREATE TABLE IF NOT EXISTS repo (
    id          INT AUTO_INCREMENT PRIMARY KEY,
    namespace   VARCHAR(32),
    name        VARCHAR(32),
    description MEDIUMBLOB,
    readme      MEDIUMBLOB,
    is_public   BOOLEAN,
    create_ts   BIGINT, 
    UNIQUE (namespace, name)
)AUTO_INCREMENT=10000;
`

var setRepoForeignKey = `
ALTER TABLE repo ADD CONSTRAINT fk__repo_account_name FOREIGN KEY(namespace) REFERENCES account(name) ON DELETE CASCADE;
`

// team
var teamTable = `
CREATE TABLE IF NOT EXISTS team (
    id          INT AUTO_INCREMENT PRIMARY KEY,
    org         VARCHAR(32),
    name        VARCHAR(32),
    description MEDIUMBLOB,
    type        VARCHAR(16),
    create_ts   BIGINT,           
    UNIQUE (org, name)
)AUTO_INCREMENT=10000;
`

var setTeamForeignKey = `
ALTER TABLE team ADD CONSTRAINT fk__team_account_name FOREIGN KEY(org) REFERENCES account(name) ON DELETE CASCADE;
`

// teammember
var teammemberTable = `
CREATE TABLE IF NOT EXISTS teammember (
    team_id INT,
    member  VARCHAR(32),
    PRIMARY KEY (team_id, member)
);
`

var setTeammemberForeignKeyTeam = `
ALTER TABLE teammember ADD CONSTRAINT fk__teammember_team_id FOREIGN KEY(team_id) REFERENCES team(id) ON DELETE CASCADE;
`

var setTeammemberForeignKeyMember = `
ALTER TABLE teammember ADD CONSTRAINT fk__teammember_account_name FOREIGN KEY(member) REFERENCES account(name) ON DELETE CASCADE;
`

// orgmember
var orgmemberTable = `
CREATE TABLE IF NOT EXISTS orgmember (
    org    VARCHAR(32),
    member VARCHAR(32),
    PRIMARY KEY (org, member)
);
`

var setOrgmemberForeignKeyOrg = `
ALTER TABLE orgmember ADD CONSTRAINT fk__orgmember_account_orgname FOREIGN KEY(org) REFERENCES account(name) ON DELETE CASCADE;
`

var setOrgmemberForeignKeyMember = `
ALTER TABLE orgmember ADD CONSTRAINT fk__orgmember_account_username FOREIGN KEY(member) REFERENCES account(name) ON DELETE CASCADE;
`

// teamaccess
var teamaccessTable = `
CREATE TABLE IF NOT EXISTS teamaccess (
    team_id INT,
    repo_id INT,
    access  VARCHAR(16),
    PRIMARY KEY (team_id, repo_id)
);
`

var setTeamaccessForeignKeyTeam = `
ALTER TABLE teamaccess ADD CONSTRAINT fk__teamaccess_team_id FOREIGN KEY(team_id) REFERENCES team(id) ON DELETE CASCADE;
`

var setTeamaccessForeignKeyRepo = `
ALTER TABLE teamaccess ADD CONSTRAINT fk__teamaccess_repo_id FOREIGN KEY(repo_id) REFERENCES repo(id) ON DELETE CASCADE;
`

// tag
var tagTable = `
CREATE TABLE IF NOT EXISTS tag (
    repo_id INT,
    name    VARCHAR(64),
    PRIMARY KEY (repo_id, name)
);
`

var setTagForeignKeyRepo = `
ALTER TABLE tag ADD CONSTRAINT fk__tag_repo_id FOREIGN KEY(repo_id) REFERENCES repo(id) ON DELETE CASCADE;
`

// hook
var hookTable = `
CREATE TABLE IF NOT EXISTS hook (
    repo_id INT,
    name    VARCHAR(32),
    url     MEDIUMBLOB,
    secret  VARCHAR(64),
    PRIMARY KEY (repo_id, name)
);
`

var setHookForeignKeyRepo = `
ALTER TABLE hook ADD CONSTRAINT fk__hook_repo_id FOREIGN KEY(repo_id) REFERENCES repo(id) ON DELETE CASCADE;
`

// Init DB
// CREATE DATABASE dhe CHARACTER SET utf8 COLLATE utf8_general_ci;

