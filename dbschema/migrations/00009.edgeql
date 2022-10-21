CREATE MIGRATION m1hyblszhkyarnsdssr5x7x6qvmjwtgxxgequoiiglm73kpnqqyouq
    ONTO m17zxmyhe4tx57vsc4ag5lo4wa4tspdhtceue2nkv25ircds4dzprq
{
  ALTER TYPE default::Avatar {
      ALTER PROPERTY avatar {
          SET default := 'https://mgjules.dev/img/avatar.webp';
      };
  };
  CREATE ABSTRACT TYPE default::Sort {
      CREATE REQUIRED PROPERTY sort -> std::int32 {
          SET default := 0;
      };
  };
  ALTER TYPE default::Links EXTENDING default::Sort LAST;
  ALTER TYPE default::Links {
      ALTER PROPERTY sort {
          RESET OPTIONALITY;
          DROP OWNED;
          RESET TYPE;
      };
  };
  CREATE TYPE default::Technologies EXTENDING default::Timestamps, default::Sort {
      CREATE REQUIRED PROPERTY name -> std::str {
          CREATE CONSTRAINT std::exclusive;
      };
  };
};
