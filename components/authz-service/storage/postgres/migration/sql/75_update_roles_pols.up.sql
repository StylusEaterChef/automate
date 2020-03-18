BEGIN;

UPDATE iam_roles
    SET
        actions = '{
            secrets:*:get,
            secrets:*:list,
            infra:nodes:get,
            infra:nodes:list,
            infra:nodeManagers:get,
            infra:nodeManagers:list,
            compliance:*:get,
            compliance:*:list,
            system:*:get,
            system:*:list,
            event:*:get,
            event:*:list,
            ingest:*:get,
            ingest:*:list,
            iam:projects:list,
            iam:projects:get
        }'
    WHERE
        id = 'viewer';

UPDATE iam_roles
    SET
        actions = '{
            infra:nodes:*,
            infra:nodeManagers:*,
            compliance:*,
            system:*,
            event:*,
            ingest:*,
            secrets:*,
            telemetry:*,
            iam:projects:list,
            iam:projects:get,
            iam:projects:assign,
            applications:*
        }'
    WHERE
        id = 'editor';

UPDATE iam_roles
    SET
        actions = '{
            infra:nodes:*,
            infra:nodeManagers:*,
            compliance:*,
            system:*,
            event:*,
            ingest:*,
            secrets:*,
            telemetry:*,
            iam:projects:list,
            iam:projects:get,
            iam:projects:assign,
            iam:policies:list,
            iam:policies:get,
            iam:policyMembers:*,
            iam:teams:list,
            iam:teams:get,
            iam:teamUsers:*,
            iam:users:get,
            iam:users:list
        }'
    WHERE
        id = 'project-owner';

COMMIT;