CREATE MIGRATION m1xrkt5u7se56m5pe475356rsacjvlhayexndr6u2y52avvzyqrlmq
    ONTO m1yvn52cmoenlpgjoidpy4xoqgfhdtr5q2fuomr5bi3yop6pjeedtq
{
  ALTER TYPE default::Links {
      CREATE REQUIRED PROPERTY sort -> std::int32 {
          SET default := 0;
      };
  };
};
