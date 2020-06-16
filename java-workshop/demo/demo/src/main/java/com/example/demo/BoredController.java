package com.example.demo;

import java.util.ArrayList;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestBody;
// import org.springframework.web.bind.annotation.RequestParam;s
import org.springframework.web.bind.annotation.RestController;

@RestController
public class BoredController {
    ArrayList<Bored> boreds = new ArrayList<>();

    @Bean
	public RestTemplate restTemplate(RestTemplateBuilder builder) {
		return builder.build();
	}

    @GetMapping("/bored")
    public Bored getBored(@RequestBody Bored bored){
        Quote quote = restTemplate.getForObject(
					"https://www.boredapi.com/api/activity", Quote.class);
        return quote;
    }
}