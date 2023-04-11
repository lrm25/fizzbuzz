import pytest

from selenium.webdriver.common.by import By
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.support.wait import WebDriverWait

from webdriver_manager.chrome import ChromeDriverManager

from test_setup import SystemTest

class TestNoBackend(SystemTest):

    def test_no_backend(self):
        
        driver = None
        try:
            driver = self.setup()
            driver.find_element(By.TAG_NAME, "button").click()
            WebDriverWait(driver, 10).until(EC.alert_is_present(), 'Alert did not appear')
        except Exception as e:
            pytest.fail(str(e))
