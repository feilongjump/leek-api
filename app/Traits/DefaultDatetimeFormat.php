<?php

namespace App\Traits;

use Carbon\CarbonInterface;

trait DefaultDatetimeFormat
{
    protected function serializeDate(\DateTimeInterface $date): string
    {
        return $date->format(CarbonInterface::DEFAULT_TO_STRING_FORMAT);
    }
}
