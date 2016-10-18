package org.kududb.examples.sample;

import org.apache.kudu.ColumnSchema;
import org.apache.kudu.Schema;
import org.apache.kudu.Type;
import org.apache.kudu.client.*;

import java.util.ArrayList;
import java.util.List;

public class Sample {

  private static final String KUDU_MASTER = System.getProperty(
      "kuduMaster", "localhost");

  public static void main(String[] args) {
    System.out.println("-----------------------------------------------");
    System.out.println("Will try to connect to Kudu master at " + KUDU_MASTER);
    System.out.println("Run with -DkuduMaster=myHost:port to override.");
    System.out.println("-----------------------------------------------");
    String tableName = "connect_test";
    KuduClient client = new KuduClient.KuduClientBuilder(KUDU_MASTER).build();

    try {
      List<ColumnSchema> columns = new ArrayList(2);
      columns.add(new ColumnSchema.ColumnSchemaBuilder("id", Type.INT32)
          .key(true)
          .build());
      columns.add(new ColumnSchema.ColumnSchemaBuilder("random_field", Type.STRING)
          .build());
      List<String> rangeKeys = new ArrayList<>();
      rangeKeys.add("id");

      Schema schema = new Schema(columns);
      client.createTable(tableName, schema,
                         new CreateTableOptions().setRangePartitionColumns(rangeKeys).setNumReplicas(1));
    } catch (Exception e) {
      e.printStackTrace();
    }
  }
}
