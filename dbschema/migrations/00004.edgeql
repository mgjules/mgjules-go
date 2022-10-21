CREATE MIGRATION m1isdq77byy5j5yhwexlrutqpsvrdldpmf4u2wzelnb6xonswadswq
    ONTO m1jvcgjkim2ybv2ch3zg4wt326nyae6b3vcnzdf344r7wcvspxaflq
{
  ALTER TYPE default::Meta {
      ALTER PROPERTY id {
          RESET OPTIONALITY;
          DROP OWNED;
          RESET TYPE;
      };
  };
};
