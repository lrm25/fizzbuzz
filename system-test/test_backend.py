import pytest

from selenium.webdriver.common.by import By
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.support.wait import WebDriverWait

from webdriver_manager.chrome import ChromeDriverManager

from test_setup import SystemTest

class clicks_equal_to_count:
    def __init__(self, clicks):
        self.clicks = clicks

    def __call__(self, driver):
        count = driver.find_element(By.CSS_SELECTOR, "div[class*='count-val']")
        return int(count.text) == self.clicks

class text_is_empty:
    
    def __call__(self, driver):
        fizz_buzz = driver.find_element(By.CSS_SELECTOR, "h1[class='fizz-buzz']")
        return fizz_buzz.text == ""
    
class TestBackend(SystemTest):

    def test_backend(self):

        clicks = 0

        try:
            driver = self.setup()
            while clicks < 17:
                driver.find_element(By.TAG_NAME, "button").click()
                clicks += 1
                WebDriverWait(driver, 2).until(clicks_equal_to_count(clicks))
                if clicks % 3 == 0:
                    WebDriverWait(driver, 1).until(EC.text_to_be_present_in_element((
                        By.CSS_SELECTOR, "h1[class='fizz-buzz']"), 'Fizz'))
                if clicks % 5 == 0:
                    WebDriverWait(driver, 1).until(EC.text_to_be_present_in_element((
                        By.CSS_SELECTOR, "h1[class='fizz-buzz']"), 'Buzz'))
                if clicks % 3 != 0 and clicks % 5 != 0:
                    WebDriverWait(driver, 1).until(text_is_empty())
        except Exception as e:
            pytest.fail(str(e))