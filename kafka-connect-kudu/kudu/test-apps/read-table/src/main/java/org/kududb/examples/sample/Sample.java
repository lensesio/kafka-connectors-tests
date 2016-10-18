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
      KuduTable table = client.openTable(tableName);
      KuduSession session = client.newSession();

      List<String> projectColumns = new ArrayList<>(1);
      projectColumns.add("random_field");
      KuduScanner scanner = client.newScannerBuilder(table)
          .setProjectedColumnNames(projectColumns)
          .build();
      while (scanner.hasMoreRows()) {
        RowResultIterator results = scanner.nextRows();
        while (results.hasNext()) {
          RowResult result = results.next();
          System.out.println(result.getString(0));
        }
      }
    } catch (Exception e) {
      e.printStackTrace();
    }
  }
}

