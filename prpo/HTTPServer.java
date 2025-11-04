import java.io.*;
import java.net.*;
import java.util.*;

class User 
{
  String username;
  String email;
  String password;

  User (String name, String email, String password)
  {
    this.username = name;
    this.email = email;
    this.password = password;
  }

  @Override
  public String toString()
  {
    return "{\"username\":\"" + this.username + "\",\"email\":\"" + this.email + "\",\"password\":\"" + this.password + "\"}";
  }
}

class HTTPServer
{
  private static HashMap<String, User> users = new HashMap<>();

  static
  {
    User n1 = new User ("Filip Dobnikar", "fd92427@student.uni-lj.si", "burek123");
    User n2 = new User ("Rok Dobnikar", "rd92113@student.uni-lj.si", "burek123221443");
    users.put("user1", n1);
    users.put("user2", n2);
  }



  public static void main (String[] args)
  {
    System.out.println("Starting server...");
    
    try (ServerSocket socket = new ServerSocket(8080))
    {
      while (true)
      {
        System.out.println("Waiting for a connection...");
        Socket clientSocket = socket.accept();
        System.out.println("Found a connection!");
        InputStream in = clientSocket.getInputStream();
        OutputStream out = clientSocket.getOutputStream();
        
        requestHandler(in, out);

        clientSocket.close();
      }
    }
    catch (IOException e)
    {
      e.printStackTrace();
    }
  } 
  private static void requestHandler (InputStream in, OutputStream out)
  {
    try
    {
      System.out.println("Handling the request...");
      BufferedReader reader = new BufferedReader(new InputStreamReader(in));

      //handle the header
      String line = reader.readLine();
      System.out.println("DEBUG: Received line: '" + line + "'");

      if (line == null)
      {
        return;
      }

      String[] lineParts = line.split(" ");

      String method = lineParts[0];
      String URI = lineParts[1];
      String version = lineParts[2];
      
      String[] path = parseURI(URI);

      if (path == null || path[0] == null || path[1] == null) return;

      PrintWriter writer = new PrintWriter(out, true);
            
      if (method.equals("GET"))
      {
        System.out.println("GET request recieved, working on it...");
        
        writer.println("HTTP/1.1 200 OK");
        writer.println();
        
        if (path.length > 1)
        {
          if (path[1].equals("users"))
          {
            if (path.length == 2)
            {
              for (String entry : users.keySet())
              {
                writer.println(users.get(entry));
              }
            }
            else 
            {
              User value = users.get(path[2]);
              if (value != null) writer.println(value);
              else writer.println("null");
            }
          }
        }

      }
      else if (method.equals("POST"))
      {
        System.out.println("POST request recieved, processing...");

        String temp; 
        
        while (!(temp = reader.readLine()).equals("{"));
        
                  
        String[] JSON = parseJSON(reader);

        User newUser = new User(JSON[0], JSON[1], JSON[2]);
        users.put(JSON[0], newUser);

        writer.println("HTTP/1.1 201 Created");
        writer.println("Content-Type: application/json");
        writer.println("Connection: close");
        writer.println();
        writer.println(newUser.toString());
        writer.flush();

        System.out.println("Successfully created a new user!");

      }
      else if (method.equals("UPDATE"))
      {
        writer.println("HTTP/1.1 200 OK");
        writer.println();
        System.out.println("UPDATE request recieved");

      }
      else if (method.equals("DELETE"))
      {
        writer.println("HTTP/1.1 200 OK");
        writer.println();
        System.out.println("DELETE request recieved");
      }
      else if (method.equals("PUT"))
      {
        writer.println("HTTP/1.1 200 OK");
        writer.println();
        System.out.println("PUT request recieved");
      }


      
    }
    catch (IOException e)
    {
      e.printStackTrace();
    }
  }

  private static String[] parseURI(String URI)
  {
    String[] path = URI.split("/");
    for (String temp : path)
    {
      System.out.println(temp);
    }
    return path;
  }
  private static String[] parseJSON(BufferedReader reader)
  {
    ArrayList<String> lines = new ArrayList<>();
    String line;
    try 
    {
      while (!(line = reader.readLine()).equals("}"))
      {
        String temp = parseLine(line);
        System.out.println(temp);
        lines.add(temp);
      }
    }
    catch (IOException e)
    {
      e.printStackTrace();
    }
    return lines.toArray(new String[0]);
  }
  private static String parseLine(String line)
  {
    String[] arr = line.split(": ");
    StringBuilder sb = new StringBuilder();
   
    for (int i = 0; i < arr[1].length(); i++)
    {
      char t = arr[1].charAt(i);
      if (t != '"' && t != ',') sb.append(t);
    }

    return sb.toString();
  }
}


