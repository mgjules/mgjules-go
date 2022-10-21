CREATE MIGRATION m1xrtckjrs66ig26nabcy4asm7gx6ucobyh4ylc37lgvhefwyk2oha
    ONTO m1xrkt5u7se56m5pe475356rsacjvlhayexndr6u2y52avvzyqrlmq
{
  ALTER TYPE default::Links {
      ALTER PROPERTY new_window {
          SET default := false;
          SET REQUIRED USING (false);
      };
  };
};
