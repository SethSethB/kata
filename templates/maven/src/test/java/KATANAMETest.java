import com.kata.KATANAME;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

public class KATANAMETest {

    KATANAME example;

    @BeforeEach
    void setup(){
        example = new KATANAME();
    }

    @Test
    void stubMethodForKATANAMEExecutes(){
        example.method();
    }

}
