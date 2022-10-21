CREATE MIGRATION m1kvff45e532kfxoqx2pmoy3cfvm5b3ro75s2ueloeofmvwcz3qtoq
    ONTO m12p3oan7mipseu55gdb7snlpky3belg2mvhuavoqeqz7ayutt5kva
{
  ALTER TYPE default::Timestamps {
      ALTER PROPERTY created_at {
          SET readonly := true;
      };
  };
};
