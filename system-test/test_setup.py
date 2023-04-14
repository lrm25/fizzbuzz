import os
from pathlib import Path
import unittest

from webdriver_manager.chrome import ChromeDriverManager

from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.common.by import By
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.support.wait import WebDriverWait

class SystemTest(unittest.TestCase):

    def setup(self):
        options = webdriver.ChromeOptions()
        options.add_argument(f"user-data-dir={Path().absolute()}/data")
        if os.getenv('GITHUB_ACTIONS') is not None:
            options.add_argument("--headless=new")
            options.add_argument('--no-sandbox')
        driver = webdriver.Chrome(service=Service(ChromeDriverManager().install()), options=options)
        driver.get('http://localhost:3000')

        WebDriverWait(driver, 10).until(
            EC.visibility_of_element_located((By.TAG_NAME, "button")),
            'Unable to locate button')
        return driver
