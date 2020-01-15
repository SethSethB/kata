package com.kata;

import com.kata.KataName;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

public class KataNameTest {

    KataName example;

    @BeforeEach
    void setup(){
        example = new KataName();
    }

    @Test
    void stubMethodForKataNameExecutes(){
        example.method();
    }

}
