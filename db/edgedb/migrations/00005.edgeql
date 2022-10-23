CREATE MIGRATION m1yvn52cmoenlpgjoidpy4xoqgfhdtr5q2fuomr5bi3yop6pjeedtq
    ONTO m1isdq77byy5j5yhwexlrutqpsvrdldpmf4u2wzelnb6xonswadswq
{
  CREATE TYPE default::Links EXTENDING default::Timestamps {
      CREATE PROPERTY alternate_url -> std::str;
      CREATE REQUIRED PROPERTY icon -> std::str;
      CREATE REQUIRED PROPERTY name -> std::str;
      CREATE PROPERTY new_window -> std::bool;
      CREATE REQUIRED PROPERTY url -> std::str;
  };
};
