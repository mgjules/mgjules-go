CREATE MIGRATION m17zxmyhe4tx57vsc4ag5lo4wa4tspdhtceue2nkv25ircds4dzprq
    ONTO m1xrtckjrs66ig26nabcy4asm7gx6ucobyh4ylc37lgvhefwyk2oha
{
  CREATE ABSTRACT TYPE default::Avatar {
      CREATE REQUIRED PROPERTY avatar -> std::str {
          SET default := 'https://avatars.githubusercontent.com/u/38979769?v=4';
      };
  };
  CREATE TYPE default::Introduction EXTENDING default::Timestamps, default::Avatar {
      CREATE REQUIRED PROPERTY introduction -> std::str;
  };
  ALTER TYPE default::Meta EXTENDING default::Avatar LAST;
  ALTER TYPE default::Meta {
      ALTER PROPERTY avatar {
          RESET OPTIONALITY;
          DROP OWNED;
          RESET TYPE;
      };
  };
};
