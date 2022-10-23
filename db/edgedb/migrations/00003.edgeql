CREATE MIGRATION m1jvcgjkim2ybv2ch3zg4wt326nyae6b3vcnzdf344r7wcvspxaflq
    ONTO m1c7q5w5cd6u57gcl6kirqji3zxfty2ypwbsxjcswjc5bgvc6hdnxq
{
  ALTER TYPE default::Meta {
      ALTER PROPERTY id {
          SET OWNED;
      };
      ALTER PROPERTY id {
          SET default := (std::uuid_generate_v4());
          SET REQUIRED;
          SET TYPE std::uuid;
      };
  };
};
